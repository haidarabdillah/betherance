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
    "enode://1cee30a102924d1ec3ef746c2bf5d6daee31983cef9e7c59939cf44d478392b34652f5d44bdf3121e2bcec1840f49ab8a1217b22ec7a87ccf79708c96a359b08@139.180.142.204:1602",
    "enode://99b02a94fb28e21e115206cdf1cf314c058112ac7b7a0ed45bc52c38ee57e110c3f37e9ec8ad2513038345cf012721c2215bc0fcdb3a8b4eb40e64f432797b11@149.28.158.174:1602",
    "enode://7e6205ef3e322ff75f7f6e40f677cd60cffda87941a3bcb2f68e16a72606c05a54c0d071a80399219c39ecfb287552ad96c1342f1d4f5470d46aaa6a443ac614@136.244.111.202:1602",
    "enode://59ad64c5db0a15e64db42b5b516889bdb2be7d230853c3b3aed97b38d2e4a19cc3cb12c07a3325a740e396be41c0f03946432a0287107f61906750e729a6782b@149.28.145.159:1602",
    "enode://53ab1f0c02d5eeee11c63f3e26fa660205a8aa2129bf8d73ab72ed860c454c9270153e189286f697a412e547c471b2908a9694e98cce9dcb6cd6998142101de7@78.141.232.216:1602",
    "enode://4034b69fae0ac2263c73fa753cee1c1136b00383bea9ad74ac457e9251ee5cdaf1aa487ce1036fd0c00685bd532e1d2e38395abd7d07288ef8134a067393f3aa@78.141.242.72:1602",
    "enode://2eeccdfa4c09f2d625201b8a367bd4c58ea861453602f9b5e3b44ca95cc1367553acb1733039e4ac14334e79b50ae251886e5461c817de75bb209087efc7f4c0@108.61.196.190:1602",
    "enode://dc1b4e390052eaec94397a062ae48f9b4e6b526b2674e84b130cd8b138f199f55bb95181f5ad6fd6f5a0975f58fd4055e8e6d04e4efbbf4fe3c40e1ce9111d2a@45.77.233.81:1602",
    "enode://0a506288571f6f315cdea92863f235e86fa7368847e828f5e8ca6f5da96c515fe9b01a7bdcf9697ef2275336dbb1f4b6fbd2343f62e1b9c301c3884359dc0016@149.28.163.81:1602",
    "enode://e31218adac1c343ceec81d8a53ff0c632632f4dd2f5b975c29282385a0bff70d37a696df69e3070e38ca35378e6fb4a5ea7f50cfe74786702f4e5fa0da8afaad@45.76.123.248:1602",
    "enode://029c59d70e1bceb3ec087318dd0794d6b47c281111955fe62972e1de0f38fc9a8371ee38b1819a51af4da4d98d3209888b6e3942533ebe8e6c9dcd677c6f5d46@45.77.245.158:1602",
    "enode://908ee23c3b18ff637f551266bb306d55d832d0bc34fa2277fa2b0f174d040febd1316d7601422482aed930e7a5f7927c980df84bfb857e4d759a1c2c754f7ffc@207.148.73.233:1602",
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
