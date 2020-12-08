package filler

import "github.com/qa_Dec_2020/models"

func FillValues(params []*models.Param, paramValue map[uint64]interface{}) {
	for _, param := range params {
		var buf []*models.Param
		buf = append(buf, param)
		for len(buf) != 0 {
			tmp := buf[0]
			buf = append(buf[:0], buf[1:]...)

			// если в values пусто
			if len(tmp.Values) == 0 {
				switch paramValue[tmp.ID].(type) {
				case string:
					tmp.Value = paramValue[tmp.ID]
				}
				continue
			}

			for _, value := range tmp.Values {
				if paramValue[tmp.ID] == uint64(value.ID) {
					tmp.Value = value.Title
				}

				buf = append(buf, value.Params...)
			}
		}
	}
}
