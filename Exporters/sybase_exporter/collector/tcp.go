package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	registerCollector("tcp", true, NewTCPCollector)
}

const (
	tcpCollectorSubsystem = "tcp"
)

var (
	tcpPacketsReceivedTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, tcpCollectorSubsystem, "tcp_packets_received_total"),
		"tcp received total",
		[]string{"tcp"}, nil,
	)
	tcpPacketsSentTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, tcpCollectorSubsystem, "tcp_packets_sent_total"),
		"tcp sent total idle",
		[]string{"tcp"}, nil,
	)
	tcpPacketsErrorrTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, tcpCollectorSubsystem, "tcp_packet_errors_total"),
		"tcp error total",
		[]string{"tcp"}, nil,
	)
	selectTCP = "SELECT @@pack_received tcp_packets_received_total, @@pack_sent tcp_packets_sent_total,@@packet_errors tcp_packet_errors_total"
)

type tcpCollector struct {
	tcpReceived *prometheus.Desc
	tcpSend     *prometheus.Desc
	tcpError    *prometheus.Desc
}

func NewTCPCollector() (Collector, error) {
	return &tcpCollector{
		tcpReceived: tcpPacketsReceivedTotal,
		tcpSend:     tcpPacketsSentTotal,
		tcpError:    tcpPacketsSentTotal,
	}, nil
}
func (c *tcpCollector) Update(ch chan<- prometheus.Metric) error {
	if err := c.updateTCP(ch); err != nil {
		return err
	}
	return nil
}
func (c *tcpCollector) updateTCP(ch chan<- prometheus.Metric) error {
	rows, err := DB.Query(selectTCP)
	if err != nil {
		return err
	}
	var receved, send, errors float64
	for rows.Next() {
		rows.Scan(&receved, &send, &errors)
	}
	ch <- prometheus.MustNewConstMetric(
		c.tcpReceived,
		prometheus.GaugeValue,
		receved,
		"packets_received_total",
	)
	ch <- prometheus.MustNewConstMetric(
		c.tcpSend,
		prometheus.GaugeValue,
		send,
		"packets_send_total",
	)
	ch <- prometheus.MustNewConstMetric(
		c.tcpError,
		prometheus.GaugeValue,
		errors,
		"packets_error_total",
	)
	return nil
}