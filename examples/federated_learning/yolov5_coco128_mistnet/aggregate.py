# Copyright 2021 The KubeEdge Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from interface import mistnet, s3_transmitter, simple_chooser
from interface import Dataset, Estimator_server
from sedna.service.server import AggregationServerV2
from sedna.common.config import BaseConfig

# from plato.models.yolo import Model
from plato.models.mindspore.ms_yolov5 import Model
from plato.config import Config

def run_server():
    data = Dataset()
    estimator = Estimator_server()

    estimator.pretrained = BaseConfig.pretrained_model_url.replace("YoloV5_for_MindSpore_0-300_274800.ckpt", "")
    estimator.saved = BaseConfig.model_url.replace("YoloV5_for_MindSpore_0-300_274800.ckpt", "")
    #estimator.model = Model('yolov5x.yaml', data.parameters["num_classes"])
    estimator.model = Model()
    server = AggregationServerV2(
        data=data,
        estimator=estimator,
        aggregation=mistnet,
        transmitter=s3_transmitter,
        chooser=simple_chooser)

    server.start()


if __name__ == '__main__':
    run_server()
