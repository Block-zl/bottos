// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  abi
 * @Author: Gong Zibin
 * @Date:   2017-01-20
 * @Last Modified by:
 * @Last Modified time:
 */

package contract

import (
	"bytes"
	"encoding/json"
)

//ABI struct for contracts
type ABI struct {
	Types   []interface{} `json:"types"`
	Structs []struct {
		Name   string            `json:"name"`
		Base   string            `json:"base"`
		Fields map[string]string `json:"fields"`
	} `json:"structs"`
	Actions []struct {
		ActionName string `json:"action_name"`
		Type       string `json:"type"`
	} `json:"actions"`
	Tables []interface{} `json:"tables"`
}

//ParseAbi parse abiraw to struct for contracts
func ParseAbi(abiRaw []byte) (*ABI, error) {
	abi := &ABI{}
	err := json.Unmarshal(abiRaw, abi)
	if err != nil {
		return &ABI{}, err
	}

	return abi, nil
}

//AbiToJson parse abi to json for contracts
func AbiToJson(abi *ABI) (string, error) {
	data, err := json.Marshal(abi)
	if err != nil {
		return "", err
	}
	return jsonFormat(data), nil
}

func jsonFormat(data []byte) string {
	var out bytes.Buffer
	json.Indent(&out, data, "", "    ")

	return string(out.Bytes())
}
