package delltdapi

import "encoding/json"

// FetchDellTD
func (c *DellTDClient) FetchWarrantyInfo(hostName string, assetTag string) (AssetInfo, error) {

	url := c.formatPath("PROD/sbil/eapi/v5/asset-entitlements?servicetags=" + assetTag)

	resp, _, _, err := c.QueryData("GET", url, nil)
	if err != nil {
		return AssetInfo{}, err
	}

	var dellInfo AssetInfo

	json.Unmarshal(resp, &dellInfo)
	if len(dellInfo) > 0 {
		dellInfo[0].HostName = hostName
	}

	return dellInfo, nil
}
