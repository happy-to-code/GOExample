package main

import "fmt"

func main() {
	var list302 = []string{"H2205000002S00005",
		"H2205000002S00007",
		"H2205000002S00009",
		"H2205000002S00008",
		"H2205000002S00012",
		"H2205000002S00010",
		"H2205000002S00011",
		"H2205000002S00013",
		"H2205000002S00014",
		"H2205000002S00015",
		"H2205000002S00016",
		"H2205000002S00017",
		"H2205000002S00018",
		"H2205000002S00019",
		"H2205000002S00020",
		"H2205000002S00021",
		"H2205000002S00022",
		"H2205000009H00023",
		"H2205000009H00025",
		"H2205000009R00024",
		"H2205000009H00026",
		"H2205000009H00028",
		"H2205000009H00029",
		"H2205000009H00030",
		"H2205000009R00031",
		"H2205000009H00032",
		"H2205000009R00033",
		"H2205000009H00034",
		"H2205000009R00035",
		"H2205000009R00036",
		"H2205000009R00037",
		"H2205000009R00038",
		"H2205000009H00039",
		"H2205000009R00040",
		"H2205000009R00041",
		"H2205000009H00042",
		"H2205000013H00043",
		"H2205000013H00044",
		"H2205000013H00045",
		"H2205000013H00046",
		"H2205000013H00047",
		"H2205000013H00051",
		"H2205000013H00052",
		"H2205000013H00053",
		"H2205000013H00055",
		"H2205000013H00056",
		"H2205000013H00057",
		"H2205000019H00058",
		"H2205000019H00059",
		"H2205000019H00060",
		"H2205000019H00062",
		"H2205000019H00063",
		"H2205000019H00064",
		"H2205000019H00065",
		"H2205000019H00066",
		"H2205000019H00068",
		"H2205000019H00067",
		"H2205000019H00069",
		"H2205000019H00070",
		"H2205000019H00071",
		"H2205000019H00072",
		"H2205000019H00073",
		"H2205000019H00074",
		"H2205000001S00076",
		"H2205000001S00075",
		"H2205000001S00077",
		"H2205000001S00078",
		"H2205000001S00079",
		"H2205000001S00080",
		"H2205000001S00081",
		"H2205000001S00082",
		"H2205000001S00083",
		"H2205000001S00084",
		"H2205000022S00085",
		"H2205000022S00086",
		"H2205000022S00087",
		"H2205000022S00088",
		"H2205000022S00089",
		"H2205000022S00091",
		"H2205000022S00090",
		"H2206000024S00319",
		"H2206000024S00320",
		"H2206000024S00321",
		"H2206000024S00322",
		"H2206000001S00323",
		"H2206000028H00324",
		"H2206000028H00325",
		"H2206000028H00326",
		"H2206000028H00327",
		"H2206000028H00328",
		"H2205000013S00049",
		"H2205000013H00048",
		"H2205000013H00050",
		"H2205000019H00061",
		"H2205000022S00092",
		"H2205000022S00093",
		"H2205000022S00094",
		"H2205000023H00095",
		"H2205000023H00097",
		"H2205000023H00096",
		"H2205000023H00099",
		"H2205000023H00100",
		"H2205000023H00101",
		"H2205000023H00102",
		"H2205000023H00103",
		"H2205000018H00106",
		"H2205000018H00107",
		"H2205000018H00108",
		"H2205000018H00109",
		"H2205000018H00110",
		"H2205000018H00111",
		"H2205000018H00112",
		"H2205000018H00113",
		"H2205000018H00114",
		"H2205000031S00116",
		"H2205000018H00115",
		"H2205000031S00117",
		"H2205000031S00119",
		"H2205000031S00118",
		"H2205000031S00120",
		"H2205000029H00121",
		"H2205000029H00122",
		"H2205000029H00123",
		"H2205000023H00098",
		"H2205000029S00124",
		"H2205000029H00125",
		"H2205000029S00126",
		"H2205000029H00128",
		"H2205000029S00127",
		"H2205000029H00130",
		"H2205000029H00129",
		"H2205000029S00131",
		"H2205000029H00132",
		"H2205000029H00133",
		"H2205000029H00135",
		"H2205000030H00137",
		"H2205000029H00134",
		"H2205000030H00138",
		"H2205000030H00136",
		"H2206000025S00152",
		"H2206000025S00150",
		"H2206000025S00151",
		"H2206000025S00153",
		"H2206000025S00154",
		"H2206000025S00155",
		"H2206000025S00156",
		"H2206000025S00157",
		"H2206000025S00158",
		"H2206000025S00159",
		"H2206000025S00160",
		"H2206000025S00161",
		"H2206000025H00163",
		"H2206000025S00162",
		"H2206000025S00164",
		"H2206000025S00165",
		"H2206000025S00166",
		"H2206000025H00167",
		"H2206000025S00169",
		"H2206000025S00168",
		"H2206000025S00170",
		"H2206000025S00171",
		"H2206000025S00172",
		"H2206000025S00173",
		"H2206000025S00174",
		"H2206000025S00175",
		"H2206000025S00176",
		"H2206000032R00183",
		"H2206000032R00185",
		"H2206000032R00184",
		"H2206000032R00187",
		"H2206000032R00186",
		"H2206000032R00188",
		"H2206000032R00189",
		"H2206000032R00190",
		"H2206000032R00191",
		"H2206000032R00192",
		"H2206000032R00193",
		"H2206000032R00194",
		"H2206000032R00196",
		"H2206000032R00195",
		"H2206000032R00197",
		"H2206000021H00198",
		"H2206000021S00200",
		"H2206000021H00201",
		"H2206000021H00202",
		"H2206000021H00203",
		"H2206000021S00204",
		"H2206000021S00205",
		"H2206000021R00206",
		"H2206000021S00207",
		"H2206000021S00208",
		"H2206000010S00209",
		"H2206000010S00210",
		"H2206000010S00211",
		"H2206000010S00212",
		"H2206000033H00213",
		"H2206000033H00214",
		"H2206000033S00215",
		"H2206000033S00216",
		"H2206000033S00217",
		"H2206000033S00218",
		"H2206000033S00219",
		"H2206000033S00220",
		"H2206000033S00222",
		"H2206000033S00223",
		"H2206000009R00224",
		"H2206000003H00228",
		"H2206000003H00229",
		"H2206000003H00230",
		"H2206000003H00231",
		"H2206000003H00232",
		"H2206000003H00233",
		"H2206000003H00234",
		"H2206000003H00235",
		"H2206000003H00236",
		"H2206000003H00237",
		"H2206000003H00238",
		"H2206000003H00239",
		"H2206000003H00240",
		"H2206000003H00242",
		"H2206000003H00241",
		"H2206000007H00243",
		"H2206000007H00244",
		"H2206000007H00245",
		"H2206000007H00246",
		"H2206000007H00247",
		"H2206000007H00248",
		"H2206000007H00249",
		"H2206000007H00250",
		"H2206000007H00251",
		"H2206000007H00252",
		"H2206000007H00253",
		"H2206000007H00254",
		"H2206000007H00255",
		"H2206000007H00256",
		"H2206000034S00258",
		"H2206000034S00259",
		"H2206000034S00260",
		"H2206000034S00261",
		"H2206000034S00262",
		"H2206000034S00263",
		"H2206000034S00264",
		"H2206000034S00265",
		"H2206000034S00266",
		"H2206000034S00267",
		"H2206000034S00268",
		"H2206000034S00269",
		"H2206000035S00270",
		"H2206000035S00271",
		"H2206000035S00272",
		"H2206000035S00273",
		"H2206000035S00274",
		"H2206000035S00275",
		"H2206000035S00276",
		"H2206000035S00277",
		"H2206000035S00278",
		"H2206000036S00280",
		"H2206000036S00279",
		"H2206000036S00281",
		"H2206000036S00282",
		"H2206000036S00283",
		"H2206000036S00284",
		"H2206000036S00285",
		"H2206000036S00286",
		"H2206000036S00287",
		"H2206000036S00288",
		"H2206000024S00290",
		"H2206000024S00289",
		"H2206000024S00291",
		"H2206000024S00292",
		"H2206000024S00293",
		"H2206000024S00294",
		"H2206000024S00295",
		"H2206000024S00296",
		"H2206000024S00297",
		"H2206000024S00299",
		"H2206000024S00300",
		"H2206000024S00298",
		"H2206000024S00301",
		"H2206000024S00302",
		"H2206000024S00304",
		"H2206000024S00303",
		"H2206000024S00305",
		"H2206000024S00306",
		"H2206000024S00307",
		"H2206000024S00309",
		"H2206000024S00310",
		"H2206000024S00308",
		"H2206000024S00311",
		"H2206000024S00312",
		"H2206000024S00313",
		"H2206000024S00314",
		"H2206000024S00315",
		"H2206000024S00316",
		"H2206000028H00329",
		"H2206000028H00330",
		"H2206000028H00331",
		"H2206000028H00332",
		"H2206000028H00333",
		"H2206000028H00334",
		"H2206000028H00335",
		"H2206000036S00336"}

	var list273 = []string{"H2205000002S00007",
		"H2205000002S00008",
		"H2205000002S00010",
		"H2205000002S00017",
		"H2205000009H00023",
		"H2205000009H00026",
		"H2205000009H00029",
		"H2205000009R00033",
		"H2205000009R00037",
		"H2205000009R00040",
		"H2205000009R00041",
		"H2205000013H00047",
		"H2205000013H00051",
		"H2205000013H00052",
		"H2205000013H00053",
		"H2205000013H00056",
		"H2205000019H00058",
		"H2205000019H00060",
		"H2205000013H00055",
		"H2205000013H00057",
		"H2205000019H00059",
		"H2205000019H00062",
		"H2205000019H00063",
		"H2205000019H00064",
		"H2205000019H00065",
		"H2205000019H00067",
		"H2205000019H00071",
		"H2205000019H00068",
		"H2205000019H00070",
		"H2205000019H00066",
		"H2205000019H00073",
		"H2205000019H00069",
		"H2205000019H00072",
		"H2205000001S00076",
		"H2205000001S00077",
		"H2205000019H00074",
		"H2205000001S00078",
		"H2205000001S00080",
		"H2205000001S00075",
		"H2205000001S00081",
		"H2205000001S00082",
		"H2205000001S00079",
		"H2205000001S00083",
		"H2205000022S00085",
		"H2205000001S00084",
		"H2205000022S00091",
		"H2205000022S00089",
		"H2205000022S00090",
		"H2205000022S00088",
		"H2205000022S00087",
		"H2205000022S00086",
		"H2206000024S00319",
		"H2206000024S00322",
		"H2206000024S00320",
		"H2206000024S00321",
		"H2206000028H00327",
		"H2206000028H00324",
		"H2206000028H00326",
		"H2206000001S00323",
		"H2206000028H00328",
		"H2206000028H00325",
		"H2205000013S00049",
		"H2205000013H00048",
		"H2205000013H00050",
		"H2205000019H00061",
		"H2205000022S00092",
		"H2205000022S00093",
		"H2205000023H00095",
		"H2205000023H00097",
		"H2205000023H00096",
		"H2205000022S00094",
		"H2205000023H00099",
		"H2205000023H00101",
		"H2205000023H00100",
		"H2205000023H00102",
		"H2205000023H00103",
		"H2205000018H00106",
		"H2205000018H00109",
		"H2205000018H00107",
		"H2205000018H00110",
		"H2205000018H00108",
		"H2205000018H00111",
		"H2205000018H00112",
		"H2205000018H00113",
		"H2205000018H00114",
		"H2205000031S00116",
		"H2205000031S00117",
		"H2205000018H00115",
		"H2205000031S00118",
		"H2205000031S00120",
		"H2205000029H00121",
		"H2205000031S00119",
		"H2205000029H00122",
		"H2205000029H00123",
		"H2205000023H00098",
		"H2205000029S00124",
		"H2205000029S00126",
		"H2205000029H00128",
		"H2205000029S00127",
		"H2205000029H00125",
		"H2205000029H00129",
		"H2205000029S00131",
		"H2206000025S00151",
		"H2205000029H00135",
		"H2206000025S00153",
		"H2205000029H00134",
		"H2206000025S00156",
		"H2206000025S00150",
		"H2206000025S00155",
		"H2205000030H00136",
		"H2206000025S00157",
		"H2206000025S00158",
		"H2206000025S00160",
		"H2206000025S00161",
		"H2206000025S00169",
		"H2206000025S00164",
		"H2206000025S00166",
		"H2206000025H00163",
		"H2206000025S00174",
		"H2206000025S00170",
		"H2206000025S00165",
		"H2206000025S00171",
		"H2206000025S00168",
		"H2206000025S00173",
		"H2206000025S00176",
		"H2206000032R00183",
		"H2206000032R00187",
		"H2206000032R00186",
		"H2206000032R00191",
		"H2206000032R00188",
		"H2206000032R00194",
		"H2206000032R00192",
		"H2206000021H00201",
		"H2206000032R00189",
		"H2206000021S00200",
		"H2206000032R00196",
		"H2206000032R00195",
		"H2206000021S00205",
		"H2206000021H00198",
		"H2206000021H00203",
		"H2206000032R00197",
		"H2206000021S00204",
		"H2206000032R00193",
		"H2206000021H00202",
		"H2206000021R00206",
		"H2206000021S00208",
		"H2206000021S00207",
		"H2206000010S00209",
		"H2206000010S00212",
		"H2206000033H00213",
		"H2206000010S00211",
		"H2206000033S00215",
		"H2206000033S00217",
		"H2206000010S00210",
		"H2206000033H00214",
		"H2206000033S00222",
		"H2206000009R00224",
		"H2206000033S00218",
		"H2206000033S00216",
		"H2206000033S00220",
		"H2206000003H00228",
		"H2206000003H00230",
		"H2206000033S00219",
		"H2206000003H00234",
		"H2206000003H00231",
		"H2206000003H00232",
		"H2206000033S00223",
		"H2206000007H00243",
		"H2206000003H00238",
		"H2206000003H00236",
		"H2206000007H00245",
		"H2206000003H00235",
		"H2206000007H00247",
		"H2206000007H00248",
		"H2206000003H00239",
		"H2206000007H00244",
		"H2206000007H00251",
		"H2206000034S00266",
		"H2206000034S00262",
		"H2206000034S00263",
		"H2206000035S00271",
		"H2206000035S00273",
		"H2206000035S00274",
		"H2206000036S00287",
		"H2206000024S00291",
		"H2206000024S00299",
		"H2206000024S00307",
		"H2206000024S00308",
		"H2206000024S00310",
		"H2206000024S00313",
		"H2206000024S00315",
		"H2206000024S00316",
		"H2206000028H00330",
		"H2206000036S00336",
		"H2206000028H00331",
		"H2206000028H00335",
		"H2206000028H00334",
		"H2206000028H00332",
		"H2206000028H00333",
		"H2206000028H00329",
		"H2205000029H00130",
		"H2205000029H00132",
		"H2205000029H00133",
		"H2205000030H00137",
		"H2205000030H00138",
		"H2206000025S00152",
		"H2206000025S00154",
		"H2206000025S00159",
		"H2206000025S00162",
		"H2206000025H00167",
		"H2206000024S00314",
		"H2206000024S00312",
		"H2206000024S00306",
		"H2206000024S00309",
		"H2206000024S00311",
		"H2206000025S00172",
		"H2206000025S00175",
		"H2206000032R00185",
		"H2206000032R00184",
		"H2206000032R00190",
		"H2206000003H00229",
		"H2206000003H00233",
		"H2206000003H00237",
		"H2206000003H00240",
		"H2206000003H00242",
		"H2206000003H00241",
		"H2206000007H00246",
		"H2206000007H00249",
		"H2206000007H00250",
		"H2206000007H00252",
		"H2206000007H00253",
		"H2206000007H00254",
		"H2206000007H00255",
		"H2206000007H00256",
		"H2206000034S00258",
		"H2206000034S00259",
		"H2206000034S00260",
		"H2206000034S00261",
		"H2206000034S00264",
		"H2206000034S00265",
		"H2206000034S00267",
		"H2206000034S00268",
		"H2206000034S00269",
		"H2206000035S00270",
		"H2206000035S00272",
		"H2206000035S00275",
		"H2206000035S00276",
		"H2206000035S00277",
		"H2206000035S00278",
		"H2206000036S00280",
		"H2206000036S00279",
		"H2206000036S00281",
		"H2206000036S00282",
		"H2206000036S00283",
		"H2206000036S00284",
		"H2206000036S00285",
		"H2206000036S00286",
		"H2206000036S00288",
		"H2206000024S00290",
		"H2206000024S00289",
		"H2206000024S00292",
		"H2206000024S00293",
		"H2206000024S00294",
		"H2206000024S00295",
		"H2206000024S00296",
		"H2206000024S00297",
		"H2206000024S00300",
		"H2206000024S00298",
		"H2206000024S00301",
		"H2206000024S00302",
		"H2206000024S00304",
		"H2206000024S00303",
		"H2206000024S00305"}

	for _, s302 := range list302 {
		exist := isExist(list273, s302)
		if !exist {
			fmt.Println(s302)
		}
	}

}
func isExist(list302 []string, l string) bool {
	for _, s302 := range list302 {
		if s302 == l {
			return true
		}
	}
	return false
}