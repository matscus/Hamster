package jmxparser

import "encoding/xml"

//JmeterTestPlan - test plan struct(head struct)
type JmeterTestPlan struct {
	XMLName    xml.Name `xml:"jmeterTestPlan"`
	Text       string   `xml:",chardata"`
	Version    string   `xml:"version,attr"`
	Properties string   `xml:"properties,attr"`
	Jmeter     string   `xml:"jmeter,attr"`
	HashTree   struct {
		Text     string   `xml:",chardata"`
		TestPlan TestPlan `xml:"TestPlan"`
		HashTree struct {
			Text                                      string                                    `xml:",chardata"`
			ComBlazemeterJmeterRandomCSVDataSetConfig ComBlazemeterJmeterRandomCSVDataSetConfig `xml:"com.blazemeter.jmeter.RandomCSVDataSetConfig"`
			HashTree                                  []struct {
				Text                  string                `xml:",chardata"`
				TransactionController TransactionController `xml:"TransactionController"`
				JSR223Sampler         []struct {
					Text       string `xml:",chardata"`
					Guiclass   string `xml:"guiclass,attr"`
					Testclass  string `xml:"testclass,attr"`
					Testname   string `xml:"testname,attr"`
					Enabled    string `xml:"enabled,attr"`
					StringProp []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
					} `xml:"stringProp"`
				} `xml:"JSR223Sampler"`
				HashTree []struct {
					Text                    string                  `xml:",chardata"`
					JSR223Sampler           []JSR223Sampler         `xml:"JSR223Sampler"`
					HashTree                []string                `xml:"hashTree"`
					ConstantThroughputTimer ConstantThroughputTimer `xml:"ConstantThroughputTimer"`
					UniformRandomTimer      UniformRandomTimer      `xml:"UniformRandomTimer"`
				} `xml:"hashTree"`
				OnceOnlyController OnceOnlyController `xml:"OnceOnlyController"`
				TestAction         TestAction         `xml:"TestAction"`
			} `xml:"hashTree"`
			Arguments                                                   []Arguments                                                   `xml:"Arguments"`
			ResultCollector                                             ResultCollector                                               `xml:"ResultCollector"`
			SetupThreadGroup                                            SetupThreadGroup                                              `xml:"SetupThreadGroup"`
			ThreadGroup                                                 []ThreadGroup                                                 `xml:"ThreadGroup"`
			ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup []ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup `xml:"com.blazemeter.jmeter.threads.concurrency.ConcurrencyThreadGroup"`
			PostThreadGroup                                             PostThreadGroup                                               `xml:"PostThreadGroup"`
		} `xml:"hashTree"`
	} `xml:"hashTree"`
}
