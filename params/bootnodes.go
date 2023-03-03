// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://eeec145d6297c980735cfe74a3b4420c91365fe1e3aaf06969051b0282127c33f6bffa26e1c3b7254d2e7d3f22d0d69de5c8b833a0f623d851c706c2ad0ed10e@95.179.194.102:31288",
	"enode://867921eca8f695d1fb70cb25b363fc58972b558cf0b49c805abcdf75ace66b789547f536b66ddc8d339015eaa845c493c5db71e8e62c82b57c3dad466b2bfb0c@192.248.169.17:31288",
	"enode://b0736d7e146c5eae960cd40e1df10ad835c451195fcf41c037641c7da19ead2650cfc802c672a60561a4771e3feb280902a015834712f2127ab5ff3f9d3b1219@45.76.136.208:31288",
	"enode://08880b8802b216f0bc345c3cf41a007120092507882568ba10b643112d43024a707fbfcffa8235fe7bb8547f714d9e89fa5b5cac30d87e205e0e1af8774a2641@45.76.144.55:31288",
	"enode://947ae581abceee4210dc3c31f43f7b37b2c6555cd78c499787f113b4d091f2afeb8d8c53b2b1e97f8b6912685593282b53ef09edf66c606bbc8fb4162faa5c9c@66.42.61.49:31288",
	"enode://2b72954810073da33da63b29f3ff463363c1582509d093d4791dfe40234afb97b2901c3ff6c71332b682dffbdc8e3750e27687164dc953aead2eb4ef90334faf@207.148.121.215:31288",
	"enode://91c62e0910d53210331e904544a0ba4c592c58528f1aebdce0bc0aa64f4d95b102a96e9adbbe93fb66480c32819bfa1b7abaaec2f396a87f1ad5e8cdb21613d7@45.32.245.231:31288",
	"enode://18f218f2391b3e69316686c7fb53fdd32652d7bebef926b86a39f99c4fdb04d5e119fdcaa17aaa5060db5f0de51f4caf3e8fb0151e944215172a91b11af62d62@45.76.121.75:31288",
	"enode://51982b67f9128621236dd2cd6c840eb98ad61018921b7ee9188ba2462ddf9776bc638f725ed618ca0b74e30f470f912b2811d90b783d9617ceff191057683449@149.28.180.82:31288",
	"enode://31a6e8e73e2c6720950b64bccc3e1d6b5f099af0d92a6ac3e9036c97c990f90e545d483772738dc70110fcffbb62d78d7a2b13912ca2153b3c25233b94b4e135@139.180.130.141:31288",
	"enode://d5ca94dce17fe1b273977ea0effdc73d382ffc64a2358e2aa250ad158b7a143e0bebcb1bce3897819bef3161b13836d2b65a4b1aeb4ce833bbe5887babeec98e@45.77.246.165:31288",
}
var TestnetBootnodes = []string{

	"enode://9e9eacf00fe30afe145f7f034eaff0ef230c8722c83bd4e49cab2ac720b1099b1bc957294cc8442797a9d6a95263e225a18b7b3308ef2ead8b0e45f7298ab7a5@34.128.125.238:31288",
	"enode://2e7063b628a102428c4b274ad12e71239fc43708a860e0ffc55174ff5d84e4cedafebd91eeae086668c558f407d276e219a39b0bcefdbe9f5d1337e6de7f287d@34.101.51.196:31288",
	"enode://eb871c92bee71f9a82864940c337dda71befd49ff9048009ad65eda36b2b98975fe9414275c85aa9154abc244f8c3a4568c05d30d98accacb4d8c007aa0f61f3@34.128.81.110:31288",
	"enode://837db97395a7e742033dc78fb29b8d9620cbe5cc48b7bae492d70c7e2efa2d2e4b647dbe5f56d68bde0f7e19fd0d9b9c1fa09cc507a4e0ba10d26a72e95576f6@34.101.54.29:31288",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
