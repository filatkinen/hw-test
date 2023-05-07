//go:build !bench
// +build !bench

package hw10programoptimization

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDomainStat(t *testing.T) {
	data := `{"Id":1,"Name":"Howard Mendoza","Username":"0Oliver","Email":"aliquid_qui_ea@Browsedrive.gov","Phone":"6-866-899-36-79","Password":"InAQJvsq","Address":"Blackbird Place 25"}
{"Id":2,"Name":"Jesse Vasquez","Username":"qRichardson","Email":"mLynch@broWsecat.com","Phone":"9-373-949-64-00","Password":"SiZLeNSGn","Address":"Fulton Hill 80"}
{"Id":3,"Name":"Clarence Olson","Username":"RachelAdams","Email":"RoseSmith@Browsecat.com","Phone":"988-48-97","Password":"71kuz3gA5w","Address":"Monterey Park 39"}
{"Id":4,"Name":"Gregory Reid","Username":"tButler","Email":"5Moore@Teklist.net","Phone":"520-04-16","Password":"r639qLNu","Address":"Sunfield Park 20"}
{"Id":5,"Name":"Janice Rose","Username":"KeithHart","Email":"nulla@Linktype.com","Phone":"146-91-01","Password":"acSBF5","Address":"Russell Trail 61"}`

	t.Run("find 'com'", func(t *testing.T) {
		result, err := GetDomainStat(bytes.NewBufferString(data), "com")
		require.NoError(t, err)
		require.Equal(t, DomainStat{
			"browsecat.com": 2,
			"linktype.com":  1,
		}, result)
	})

	t.Run("find 'gov'", func(t *testing.T) {
		result, err := GetDomainStat(bytes.NewBufferString(data), "gov")
		require.NoError(t, err)
		require.Equal(t, DomainStat{"browsedrive.gov": 1}, result)
	})

	t.Run("find 'unknown'", func(t *testing.T) {
		result, err := GetDomainStat(bytes.NewBufferString(data), "unknown")
		require.NoError(t, err)
		require.Equal(t, DomainStat{}, result)
	})
}

func BenchmarkGetDomainStatRecords(b *testing.B) {

	data := `{"Id":1,"Name":"Howard Mendoza","Username":"0Oliver","Email":"aliquid_qui_ea@Browsedrive.gov","Phone":"6-866-899-36-79","Password":"InAQJvsq","Address":"Blackbird Place 25"}
{"Id":2,"Name":"Jesse Vasquez","Username":"qRichardson","Email":"mLynch@broWsecat.com","Phone":"9-373-949-64-00","Password":"SiZLeNSGn","Address":"Fulton Hill 80"}
{"Id":3,"Name":"Clarence Olson","Username":"RachelAdams","Email":"RoseSmith@Browsecat.com","Phone":"988-48-97","Password":"71kuz3gA5w","Address":"Monterey Park 39"}
{"Id":4,"Name":"Gregory Reid","Username":"tButler","Email":"5Moore@Teklist.net","Phone":"520-04-16","Password":"r639qLNu","Address":"Sunfield Park 20"}
{"Id":5,"Name":"Janice Rose","Username":"KeithHart","Email":"nulla@Linktype.com","Phone":"146-91-01","Password":"acSBF5","Address":"Russell Trail 61"}`

	for i := 0; i < b.N; i++ {
		result, _ := GetDomainStat(bytes.NewBufferString(data), "com")
		_ = result
	}
}

func TestGetDomainStatAdd(t *testing.T) {
	var data = `{"Id":186,"Name":"Russell Campbell","Username":"DavidMatthews","Email":"voluptatum_odit_rerum@Dabshots.info","Phone":"0-501-398-20-95","Password":"KMbwO6DtMl","Address":"Northview Drive 12"}
{"Id":187,"Name":"Chris Hicks","Username":"3Fields","Email":"wRobinson@Shuffledrive.gov","Phone":"139-97-32","Password":"zKTjbv1Nk","Address":"Redwing Junction 39"}
{"Id":188,"Name":"Christine Williamson","Username":"BillyLee","Email":"LawrenceScott@Voonyx.biz","Phone":"6-412-508-26-10","Password":"lsYP4y6PvA","Address":"Melvin Road 56"}
{"Id":189,"Name":"Melissa Alvarez","Username":"animi_harum","Email":"sit@Feedspan.net","Phone":"721-42-09","Password":"gTETQDkwP1o","Address":"Iowa Lane 78"}
{"Id":190,"Name":"Jason Fernandez","Username":"accusantium_laudantium_et","Email":"RussellWood@Jamia.com","Phone":"3-952-117-47-78","Password":"EUmHSrkvqmB6","Address":"Miller Point 79"}
{"Id":191,"Name":"Mr. Dr. Walter Dean","Username":"et","Email":"HarryKing@Oyoloo.info","Phone":"9-633-583-99-70","Password":"9a8erb","Address":"Ilene Park 82"}
{"Id":192,"Name":"Philip Ruiz","Username":"aHansen","Email":"ChristinaJordan@Riffpath.net","Phone":"752-10-60","Password":"w2S0fi","Address":"Bowman Court 92"}
{"Id":193,"Name":"Brian Sims","Username":"fDixon","Email":"lDixon@Zoomcast.name","Phone":"6-653-049-59-29","Password":"3wd2Fr","Address":"Blue Bill Park Junction 77"}
{"Id":194,"Name":"Steve Castillo","Username":"vArmstrong","Email":"ab_facilis@Blogpad.org","Phone":"9-967-484-33-43","Password":"zM7zRPU","Address":"Russell Crossing 11"}
{"Id":195,"Name":"Laura Porter I II III IV V MD DDS PhD DVM","Username":"consequatur_cumque","Email":"1Barnes@Zoombeat.name","Phone":"3-414-774-33-02","Password":"IpAFVJwtM","Address":"Sutteridge Street 3"}
{"Id":196,"Name":"Juan Medina","Username":"GaryHunt","Email":"facilis_asperiores@Jabberbean.edu","Phone":"4-692-174-70-77","Password":"1X3iv92mR","Address":"Jay Center 98"}
{"Id":197,"Name":"Jack Rice","Username":"CraigBurns","Email":"JuanLopez@Kayveo.biz","Phone":"7-481-140-09-92","Password":"XNpJnSWn","Address":"Killdeer Lane 68"}
{"Id":198,"Name":"Diane Duncan","Username":"AliceRay","Email":"7Howell@Rhyloo.com","Phone":"0-526-790-68-97","Password":"2OtQCFw","Address":"Shoshone Parkway 27"}
{"Id":199,"Name":"Mrs. Ms. Miss Virginia Snyder","Username":"sFord","Email":"zMcdonald@Aimbo.edu","Phone":"7-881-230-06-92","Password":"AumLxsJRByeV","Address":"High Crossing Park 85"}
{"Id":200,"Name":"Juan Peters","Username":"soluta","Email":"aut_explicabo_voluptatum@Cogibox.com","Phone":"4-732-469-97-06","Password":"hqTPJCx4kCdD","Address":"Moose Avenue 3"}`

	t.Run("find 'biz'", func(t *testing.T) {
		result, err := GetDomainStat(bytes.NewBufferString(data), "biz")
		require.NoError(t, err)
		require.Equal(t, DomainStat{"kayveo.biz": 1, "voonyx.biz": 1}, result)
	})
}
