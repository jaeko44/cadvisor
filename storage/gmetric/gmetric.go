/* 
** Copyright [2013-2015] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
*/

package gmetricc

import (
	"fmt"
	"sync"
	"time"
	"github.com/facebookgo/ganglia/gmetric"	
	info "github.com/google/cadvisor/info/v1"
//	"github.com/jbuchbinder/go-gmetric/gmetric"
	"net"
)

type gmetricc struct {	
	lock           sync.Mutex
}

const (
	colTimestamp          string = "time"
	colMachineName        string = "machine"
	colContainerName      string = "container_name"
	colCpuCumulativeUsage string = "cpu_cumulative_usage"
	// Memory Usage
	colMemoryUsage string = "memory_usage"
	// Working set size
	colMemoryWorkingSet string = "memory_working_set"
	// Cumulative count of bytes received.
	colRxBytes string = "rx_bytes"
	// Cumulative count of receive errors encountered.
	colRxErrors string = "rx_errors"
	// Cumulative count of bytes transmitted.
	colTxBytes string = "tx_bytes"
	// Cumulative count of transmit errors encountered.
	colTxErrors string = "tx_errors"
	// Filesystem device.
	colFsDevice = "fs_device"
	// Filesystem limit.
	colFsLimit = "fs_limit"
	// Filesystem usage.
	colFsUsage = "fs_usage"
)

func (self *gmetricc) AddStats(ref info.ContainerReference, stats *info.ContainerStats) error {
	if stats == nil {
		return nil
	}
	fmt.Println("----------------addstats--------------------")
	
	//ab := gmetric.ClientFromFlag("hello")
	ab := gmetric.Client{
			Addr: []net.Addr{
				&net.UDPAddr{IP: net.ParseIP("103.56.92.26"), Port: 8651},
		     },
			}

	m := &gmetric.Metric{
		Name:         "uint8_metric",
		Host:         "103.56.92.26",
		ValueType:    gmetric.ValueUint8,
		Units:        "count",
		Slope:        gmetric.SlopeBoth,
		TickInterval: 20 * time.Second,
		Lifetime:     24 * time.Hour,
	}
	const val = 10

	if err1 := ab.WriteMeta(m); err1 != nil {
		fmt.Println(err1)
	}

	if err2 := ab.WriteValue(m, val); err2 != nil {
		fmt.Println(err2)
	}
	
	

	
		return nil
}

func (self *gmetricc) Close() error {
	//self.client = nil
	return nil
}

// machineName: A unique identifier to identify the host that current cAdvisor
// instance is running on.
// influxdbHost: The host which runs influxdb.
func New() (*gmetricc, error) {
	fmt.Println("================================gmetric entry==========")
	ret := &gmetricc{}
	return ret, nil
}
