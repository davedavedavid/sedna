package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"ct_exporter/api/types/v1alpha1"
	clientV1alpha1 "ct_exporter/clientset/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type TrainReadyData struct {
	Input struct {
		Models []struct {
			Format  string   `json:"format"`
			URL     string   `json:"url"`
			Devices []string `json:"devices"`
		} `json:"models"`
		DataURL string       `json:"dataURL"`
		DataIndexURL string  `json:"dataIndexURL"`
		OutputDir   string   `json:"outputDir"`
	} `json:"input"`
}
type TrainCompletedData struct {
	Output struct {
		Models []struct {
			Format  string   `json:"format"`
			URL     string   `json:"url"`
			Devices []string `json:"devices"`
		} `json:"models"`
	} `json:"output"`
}


var kubeconfig string

var (
	TrainProb          *prometheus.Desc
	Status             *prometheus.Desc
	Stage              *prometheus.Desc
	FullStage          *prometheus.Desc
	NumberOfSamples    *prometheus.Desc
	LastHeartBreakTime *prometheus.Desc
	TrainModelPath     *prometheus.Desc
	EvalNewModelMetric *prometheus.Desc
	EvalOldModelMetric *prometheus.Desc
	DeployModelPath    *prometheus.Desc
	IlVersion          *prometheus.Desc
)

type  Exporter struct {
	client    *clientV1alpha1.ExampleV1Alpha1Client
}

func NewExporter() (*Exporter, error) {
	var config *rest.Config
	var err error

	if kubeconfig == "" {
		log.Printf("using in-cluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		log.Printf("using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		panic(err)
	}

	err = v1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		panic(err)
	}
	clientSet, err := clientV1alpha1.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return &Exporter{
		client: clientSet,
	}, nil
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- TrainProb
	ch <- Status
	ch <- NumberOfSamples
	ch <- Stage
	ch <- FullStage

	ch <- LastHeartBeatTime
	ch <- LastProbeTime

	ch <- TrainModelPath
	ch <- EvalNewModelMetric
	ch <- EvalOldModelMetric
	ch <- DeployModelPath
	ch <- IlVersion
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	ctx := context.Background()
	federatedRes, err := e.client.FederatedClient("default").Get(ctx, "ct-yolo-v5", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	//trainProb := federatedRes.Spec.Dataset.TrainProb
	//ch <- prometheus.MustNewConstMetric(
	//	TrainProb,
	//	prometheus.GaugeValue,
	//	trainProb)

	conditions := federatedRes.Status.Conditions
	//status := conditions[len(conditions)-1].Status

	//fakemetric?
	fakeMetric := float64(rand.Intn(10))
	ch <- prometheus.MustNewConstMetric(
		Status,
		prometheus.GaugeValue,
		fakeMetric)

	stage := conditions[len(conditions)-1].Stage
	stageType := conditions[len(conditions)-1].Type
	fullStage := string(stage) + string(stageType)
	ch <- prometheus.MustNewConstMetric(
		FullStage,
		prometheus.GaugeValue,
		float64(stageMap[fullStage]))

	ch <- prometheus.MustNewConstMetric(
		Stage,
		prometheus.GaugeValue,
		float64(stageMap[string(stage)]))

	lastHeartBeatTime := conditions[len(conditions)-1].LastHeartbeatTime
	heartBeatTimeUnix := lastHeartBeatTime.Unix()
	ch <- prometheus.MustNewConstMetric(
		LastHeartBreakTime,
		prometheus.GaugeValue,
		float64(heartBeatTimeUnix)*1000)

	lastProbeTime := conditions[len(conditions)-1].LastProbeTime
	probeTimeUnix := lastProbeTime.Unix()
	ch <- prometheus.MustNewConstMetric(
		lastProbeTime,
		prometheus.GaugeValue,
		float64(probeTimeUnix)*1000)

	trainModelPath                := "null"
	newEvalModelPath              := "null"
	newEvalModelMap               := "null"
	newEvalModelAllPrecision      := "null"
	newEvalModelPrecision         := "null"
	newEvalModelAllRecall         := "null"
	newEvalModelRecall            := "null"
	oldEvalModelPath              := "null"
	oldEvalModelMap               := "null"
	oldEvalModelAllPrecision      := "null"
	oldEvalModelPrecision         := "null"
	oldEvalModelAllRecall         := "null"
	oldEvalModelRecall            := "null"
	deployModelPath               := "null"
	deployStatus                  := "waiting"
	ilVersion                     := 0.0


	var condition v1alpha1.ILJobCondition
	var conditionStage string
	var rawData string
	var stringValue []string

	for i := len(conditions) - 1; i >= 0; i-- {
		condition = conditions[i]
		conditionStage = string(condition.Stage) + string(condition.Type)
		rawData = condition.Data

		if i != len(conditions) - 1 && conditionStage == "TrainWaiting" {
			break
		}
		/*
		if i == len(conditions) - 1 && conditionStage == "TrainReady" {
			deployStatus = "waiting"
		}
		*/
		switch conditionStage {
			case "TrainReady":
				var trainReadyData TrainReadyData
				err := json.Unmarshal([]byte(rawData), &trainReadyData)
				if err != nil {
					panic(err)
				}
				outputDir := trainReadyData.Input.OutputDir
				fmt.Printf("%+v\n", outputDir)
				ilVersion, _ = strconv.ParseFloat(filepath.Base(outputDir), 64)

			case "TrainCompleted":
				var trainCompletedData TrainCompletedData
				err := json.Unmarshal([]byte(rawData), &trainCompletedData)
				if err != nil {
					panic(err)
				}
				lenTrainModelPaths := len(trainCompletedData.Output.Models)
				// 保证ckpt路径是最后一个
				trainModelPath = trainCompletedData.Output.Models[lenTrainModelPaths-1].URL

			case "EvalReady":
				var evalReadyData EvalReadyData
				err := json.Unmarshal([]byte(rawData), &evalReadyData)
				if err != nil {
					panic(err)
				}
				newEvalModelPath = evalReadyData.Input.Models[0].URL
				oldEvalModelPath = evalReadyData.Input.Models[1].URL

			case "EvalCompleted":
				var evalCompletedData EvalCompletedData
				err := json.Unmarshal([]byte(rawData), &evalCompletedData)
				if err != nil {
					panic(err)
				}

				newModelMetric := evalCompletedData.Output.Models[0].Metrics
				for metric, values := range newModelMetric{
					stringValue = []string{}
					for _, value := range values {

						stringValue = append(stringValue, strconv.FormatFloat(value, 'f', 3, 64))
					}
					switch metric {
					case "mAP": newEvalModelMap = strings.Join(stringValue, ",")
					case "all_precision": newEvalModelAllPrecision = strings.Join(stringValue, ",")
					case "all_recall": newEvalModelAllRecall = strings.Join(stringValue, ",")
					case "precision": newEvalModelPrecision = strings.Join(stringValue, ",")
					case "recall": newEvalModelRecall = strings.Join(stringValue, ",")
					}
				}

				oldModelMetric := evalCompletedData.Output.Models[1].Metrics
				for metric, values := range oldModelMetric{
					stringValue = []string{}
					for _, value := range values {
						stringValue = append(stringValue, strconv.FormatFloat(value, 'f', 3,64))
					}
					switch metric {
					case "mAP": oldEvalModelMap = strings.Join(stringValue, ",")
					case "all_precision": oldEvalModelAllPrecision = strings.Join(stringValue, ",")
					case "all_recall": oldEvalModelAllRecall = strings.Join(stringValue, ",")
					case "precision": oldEvalModelPrecision = strings.Join(stringValue, ",")
					case "recall": oldEvalModelRecall = strings.Join(stringValue, ",")
					}
				}
			case "DeployReady":
				var deployReadyData DeployReadyData
				err := json.Unmarshal([]byte(rawData), &deployReadyData)
				if err != nil {
					panic(err)
				}
				deployModelPath = deployReadyData.Input.Models[0].URL
				//stageCondition := conditions[len(conditions)-1].Type
			case "DeployCompleted":
				fmt.Printf(rawData)
				if rawData == "{}" {
					deployStatus = "False"
				} else {
					deployStatus = "True"
				}
		}
	}

	ch <- prometheus.MustNewConstMetric(
		TrainModelPath,
		prometheus.GaugeValue,
		fakeMetric,
		trainModelPath)
	ch <- prometheus.MustNewConstMetric(
		EvalNewModelMetric,
		prometheus.GaugeValue,
		fakeMetric,
		newEvalModelPath,
		newEvalModelMap,
		newEvalModelAllPrecision,
		newEvalModelAllRecall,
		newEvalModelPrecision,
		newEvalModelRecall)
	ch <- prometheus.MustNewConstMetric(
		EvalOldModelMetric,
		prometheus.GaugeValue,
		fakeMetric,
		oldEvalModelPath,
		oldEvalModelMap,
		oldEvalModelAllPrecision,
		oldEvalModelAllRecall,
		oldEvalModelPrecision,
		oldEvalModelRecall)
	ch <- prometheus.MustNewConstMetric(
		DeployModelPath,
		prometheus.GaugeValue,
		fakeMetric,
		deployModelPath,
		deployStatus)
	ch <- prometheus.MustNewConstMetric(
		IlVersion,
		prometheus.GaugeValue,
		ilVersion)

	datasetRes, err := e.client.DatasetClient("default").Get(ctx, "federated-dataset", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	number := datasetRes.Status.NumberOfSamples
	ch <- prometheus.MustNewConstMetric(
		NumberOfSamples,
		prometheus.GaugeValue,
		float64(number))
}


func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.Parse()
}

func main() {
	TrainProb = prometheus.NewDesc(
		"inclearning_train_prob",
		"train_prob of dataset",
		nil,
		nil)
	Status = prometheus.NewDesc(
		"inclearning_status",
		"status",
		nil,
		nil)
	Stage = prometheus.NewDesc(
		"inclearning_stage",
		"stage",
		nil,
		nil)
	FullStage = prometheus.NewDesc(
		"inclearning_full_stage",
		"full stage",
		nil,
		nil)
	NumberOfSamples = prometheus.NewDesc(
		"number_of_samples",
		"number_of_samples",
		nil,
		nil)
	LastHeartBreakTime = prometheus.NewDesc(
		"LastHeartBreakTime",
		"LastHeartBreakTime",
		nil,
		nil)
	TrainModelPath = prometheus.NewDesc(
		"TrainModelPath",
		"TrainModelPath",
		[]string{"train_model_path"},
		nil)
	EvalNewModelMetric = prometheus.NewDesc(
		"EvalNewModelMetric",
		"EvalNewModelMetric",
		[]string{"a_eval_new_model_path", "a_mAP", "all_precision",
			"all_recall", "precision", "recall"},
		nil)
	EvalOldModelMetric = prometheus.NewDesc(
		"EvalOldModelMetric",
		"EvalOldModelMetric",
		[]string{"a_eval_old_model_path", "a_mAP", "all_precision",
			"all_recall", "precision", "recall"},
		nil)
	DeployModelPath = prometheus.NewDesc(
		"DeployModelPath",
		"DeployModelPath",
		[]string{"deploy_model_path", "deploy_status"},
		nil)
	IlVersion = prometheus.NewDesc(
		"IlVersion",
		"IlVersion",
		nil,
		nil)
	exporter, err := NewExporter()
	if err != nil {
		panic(err)
	}
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(exporter)

	gatherers := prometheus.Gatherers{
		prometheus.DefaultGatherer,
		reg,
	}
	h := promhttp.HandlerFor(gatherers,
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
		})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	fmt.Printf("Start server at :9104")
	if err := http.ListenAndServe(":9104", nil); err != nil {
		fmt.Printf("Error occur when start server %v", err)
	}

}
