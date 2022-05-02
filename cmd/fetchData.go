package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

const VERSION_URL = "https://ddragon.leagueoflegends.com/api/versions.json"

func init() {
	rootCmd.AddCommand(fetchDataCommand)
}

func getFetchCommand(url string) string {
	return fmt.Sprintf("curl %v | tar -xz", url)
}

func getFetchUrl(version string) string {
	return fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/dragontail-%v.tgz", version)
}

var fetchDataCommand = &cobra.Command{
	Use:   "fetch_cmd",
	Short: "Prints command to fetch last champ data",
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient := http.Client{}
		req, err := http.NewRequest(http.MethodGet, VERSION_URL, nil)
		if err != nil {
			return err
		}
		res, err := httpClient.Do(req)
		if err != nil {
			return err
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		var versions []string
		err = json.Unmarshal(body, &versions)
		if err != nil {
			return err
		}

		fetchUrl := getFetchUrl(versions[0])
		fileRes, err := httpClient.Head(fetchUrl)
		if err != nil {
			return err
		}
		contLen := fileRes.Header.Get("Content-Length")
		var gbSize float64
		if contLen != "" {
			byteSize, err := strconv.ParseInt(contLen, 10, 32)
			if err == nil {
				gbSize = float64(byteSize) / 1073741824
			}
		}
		fmt.Printf(
			"Champions archive size is %v GB\nUse following command to download it:\n\n	%v\n",
			math.Floor(gbSize*100)/100,
			getFetchCommand(fetchUrl),
		)
		return nil
	},
}
