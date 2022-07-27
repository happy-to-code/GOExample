package main

import (
	"encoding/json"
	"fmt"
)

type Cjrff struct {
	Code int `json:"code"`
	Data []struct {
		CONPHONE           string `json:"CONPHONE"`
		CERTTIME           string `json:"certTime"`
		EYELEVEL           string `json:"eyeLevel"`
		IDCARD             string `json:"idCard"`
		BRITHTIME          string `json:"brithTime"`
		SEX                string `json:"sex"`
		STATE              int    `json:"state"`
		EDUCATION          string `json:"education"`
		FIRSTTIME          string `json:"firstTime"`
		AREACODE           string `json:"areaCode"`
		KINDSTR            string `json:"kindStr"`
		CJRCJRLX           string `json:"cjrCjrlx"`
		NATION             string `json:"nation"`
		DISABLEDNUMBER     string `json:"disabledNumber"`
		EARLEVEL           string `json:"earLevel"`
		EVALUATE           string `json:"evaluate"`
		DISABLEDLEVEL      string `json:"disabledLevel"`
		BODYLEVEL          string `json:"bodyLevel"`
		RESIDENCEAREA      string `json:"residenceArea"`
		STIME              string `json:"sTime"`
		ZXREASON           string `json:"zxReason"`
		ETIME              string `json:"etime"`
		DISABLEDTYPE       string `json:"disabledType"`
		GUARDNAME          string `json:"guardName"`
		SPEECHLEVEL        string `json:"speechLevel"`
		NAME               string `json:"name"`
		DEFORMITYTYPE      string `json:"deformityType"`
		CJRALLOCATIONSTATE string `json:"cjrAllocationState"`
		UTIME              string `json:"uTime"`
		IQLEVEL            string `json:"iqLevel"`
		MENTALLEVEL        string `json:"mentalLevel"`
		PROVINCEID         int    `json:"provinceId"`
	} `json:"data"`
	Message string `json:"message"`
}

func main() {
	var s = `{"code":200,"data":[{"CON_PHONE":"13222318768","CERTTIME":"2020-08-13 00:00:00","EYE_LEVEL":"","IDCARD":"321081197109121224","BRITH_TIME":"1971-09-12 00:00:00","SEX":"2","STATE":1,"EDUCATION":"3","FIRST_TIME":"2009-11-12 00:00:00","AREA_CODE":"321081401207","KINDSTR":"5,","CJR_CJRLX":"0","NATION":"1","DISABLED_NUMBER":"32108119710912122454","EAR_LEVEL":"","EVALUATE":"同意换证","DISABLED_LEVEL":"4","BODY_LEVEL":"","RESIDENCE_AREA":"江苏省扬州市仪征市枣林湾旅游度假区长山村委会钱长组","STIME":"2020-08-13 00:00:00","ZX_REASON":"","ETIME":"2030-08-13 00:00:00","DISABLED_TYPE":"5","GUARD_NAME":"钱修林","SPEECH_LEVEL":"","NAME":"王华英","DEFORMITY_TYPE":"","CJR_ALLOCATION_STATE":"0","UTIME":"2021-05-15 00:00:00","IQ_LEVEL":"4","MENTAL_LEVEL":"","PROVINCEID":32}],"message":""}`
	var cjr Cjrff2
	err := json.Unmarshal([]byte(s), &cjr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cjr)
	bytes, _ := json.Marshal(cjr)
	fmt.Printf("%s\n", bytes)
}

type Cjrff2 struct {
	Code int `json:"code"`
	Data []struct {
		CON_PHONE            string `json:"con_phone"`
		CERTTIME             string `json:"certTime"`
		EYE_LEVEL            string `json:"eye_level"`
		IDCARD               string `json:"idCard"`
		BRITH_TIME           string `json:"brith_time"`
		SEX                  string `json:"sex"`
		STATE                int    `json:"state"`
		EDUCATION            string `json:"education"`
		FIRST_TIME           string `json:"first_time"`
		AREA_CODE            string `json:"area_code"`
		KINDSTR              string `json:"kindStr"`
		CJR_CJRLX            string `json:"cjr_cjrlx"`
		NATION               string `json:"nation"`
		DISABLED_NUMBER      string `json:"disabled_number"`
		EAR_LEVEL            string `json:"ear_level"`
		EVALUATE             string `json:"evaluate"`
		DISABLED_LEVEL       string `json:"disabled_level"`
		BODY_LEVEL           string `json:"body_level"`
		RESIDENCE_AREA       string `json:"residence_area"`
		STIME                string `json:"sTime"`
		ZX_REASON            string `json:"zx_reason"`
		ETIME                string `json:"etime"`
		DISABLED_TYPE        string `json:"disabled_type"`
		GUARD_NAME           string `json:"guard_name"`
		SPEECH_LEVEL         string `json:"speech_level"`
		NAME                 string `json:"name"`
		DEFORMITY_TYPE       string `json:"deformity_type"`
		CJR_ALLOCATION_STATE string `json:"cjr_allocation_state"`
		UTIME                string `json:"uTime"`
		IQ_LEVEL             string `json:"iq_level"`
		MENTAL_LEVEL         string `json:"mental_level"`
		PROVINCEID           int    `json:"provinceId"`
	} `json:"data"`
	Message string `json:"message"`
}
