// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package nats

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "nats", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzUl8FS4zgQhu95ii5Ouwd4gBy2amv3wmE4DMw5KHLH0YytNq1WqPD0U1Ji49hy4gRDZXzgEIn+P/9S/5Jv4Rdu52CVuBmAGClwDjcPStzNDCBDp9lUYsjO4Z8ZAMSZ8I0yX+AMYGWwyNw8jtyCVSU2tcIj2wrnkDP5av9LomJ4nsM/PYMmK8pYB06UGCdGO5C1EnhFRmBUGayYSnh4l2gTtCkc8gZ5YbJmpMZxwsbmrZ8HmMLztMZ9Jbj/v6dh6bVXPVOC42o/7uqKKRFoBSUKGw2aUYXZPTFN1qIOQ64n2nb4hOp/tcVxId+LAmOhBLM9R1uja3GbSkhUcTBSMxkrmCN3xo6Q1X5bXy6RgyXaM6OVYgtKi9kg6MKgFdfzhskLTmhLrHeNjjDmxgkyZt13bva9KJnQiY3iC33wVdjZSSMKOui/kS5UyIYykPeWNA58BX851H8nEUos75bbw53xUYr9ngy1ibfgncpj+z78+/QIFZNG55IwmngAZIJ9UVButCp2ItGiNg+wtw5amXLAVfkk1aogJZf7oyt/YI47ak5smkU64iY1KgoV213MYtZLlDYUY0lDm2fiXk4o1RTGJgG6HT1SXpXkrQR5YzWVxubhxFKduam+hoPGcmFpu9acsmcE4xHOQdGaKtXqn4vUV6xZyHd7Z6JlIy85/QHL1nBez7I1SMPL5gp6DTHkfIn8aZ0fVKBRaV1zyMbcTqKtRaoF48sinvB3TCQLzybJSMufqM8M7w7j2kg4MiDoQEnWCIVrM3QlB/BClL9NypfGizoX8MXb01cA7oQuIHR+6b4CMOpcwhf+fgVf0BnLN2Xmhnr1nbe54e1unmfmbmPaYMqlEEdg1qiMLx6dtL+bE3OHMNuo/fUcBzoStgMMmQlf40sfv0DJtpf5x/f7x4Eax14Eup+Fb4PTTh8xZ74ZfCAFUvwx3K4B/2TKpug3iq8CPnCcyx5D8RrgT6Zzij6c2dcAn7o7/A4AAP//3EU44A=="
}
