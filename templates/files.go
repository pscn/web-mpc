// Code generated by "file2go -output templates/files.go templates/index.html templates/javascript.js templates/style.css templates/theme-default.css templates/theme-juri.css"; DO NOT EDIT.

// Encoded files:
// → templates/index.html
// → templates/javascript.js
// → templates/style.css
// → templates/theme-default.css
// → templates/theme-juri.css

package templates

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"time"
)

// file stores the content and modTime of the decoded string
type file struct {
	content *[]byte
	modTime *time.Time
}

var container map[string]*file

// Init populates *file with data decoded from base64Encoded string
func decode(base64Encoded string) (*file, error) {
  gzipEncoded, err := base64.StdEncoding.DecodeString(base64Encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data(BASE64): %s", err)
	}

	var buf bytes.Buffer
	_, err = buf.Write(gzipEncoded)
	if err != nil {
		return nil, fmt.Errorf("failed buffer decode data: %s", err)
	}

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create reader from buffer: %s", err)
	}
	defer zr.Close()

  decoded, err := ioutil.ReadAll(zr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data(GZIP): %s", err)
	}

	return &file{
		content: &decoded,
		modTime: &zr.ModTime,
	}, nil
}

func init() {
  container = make(map[string]*file)
  var err error
  
  container["templates/index.html"], err = decode(`` +
    `H4sICFKhsVwA/2FXNWtaWGd1YUhSdGJBPT0A7Btdj+O28T2/YqK+NIgki5JlWwtri8` +
    `vdtSmQSxd7h+3HGy3RNhNJVEnaXifIfy9IyrJky18b7+5dEaDoHqnhfM9wOBOPv373` +
    `j7ef/n33HuYyz26/Gqs/kOFiFluksG6/AhjPCU7VPwDGOZEYkjnmgsjYWsipM7Kanw` +
    `qck9jK2YRmxFmRiYPL0klwiScZsSBhhSSFjK01ER3HlpSsSsZlA3BFUzmPU7KkCXH0` +
    `wgZaUElx5ogEZyRGtsYDIOacFj87kjlTKuOCbQhktPgZOMliS8h1RsScEGnBnJNpte` +
    `MmQpwDLOckJ05KpniRyeahrx2ncdoAz6UsxU2vN2WFFO6MsVlGcEmFm7C8lwjxlynO` +
    `abaO79mESfbtB1YwqxJkl/xtte84t08nQ2dzSdhCHCcyllRm5PafZAIf7t6Oe2apfK` +
    `C3cYLxhKVroGlsqX80VAAlFgJKzHFOJOECppzlkGNauDMGksGcpikpgBblQsKUkiwV` +
    `W5lSuoQkw0LE1pympBZ6bMAVvZWwQK5LEluSPEoLljhbkNj69Vd3JX77bcNJL6XL26` +
    `+2XH18+BvQhBWGltkXyxlgTrFjWIotyRfEgilLFkJ5amxNcSaIBVpBsZVSUWZ4fVOw` +
    `osGZWOcTlmnWFAEnoTxRXq7c+Dv2GFseeOCr/9VnAMYGCpLH2EKeBcna/OWxNbJuxz` +
    `3zuabRM0QO0lRsnaJYYjmHNLY+QGAH8ABoCD8AGtrIg/8okur7+QTxQpyUkZNEggld` +
    `q2/BnCjviy3Ut+BR76xjK1CkFeDZ55DfffAkz4VxlzOV5G+VFHXqqAWNUEOnoyfqlJ` +
    `Pl+fyhUYMiOs1htAX3n8afkKy8xOSjrelG2nIDbbjBxYYTBPNkfr5qlL3cEAIbDTtU` +
    `0oy8wETeUAde6IZPCL2Migu8SrnJ99CHBwggOGqvgYZEIwV63PXsqEIZQXQCZbRBeR` +
    `wQ2SiscKIQUHgCqwbWaPdBT+pvyrKU8HM1mINv+zCHEDIIbN8NYQ6a8Ai+B/8pQcdW` +
    `p6nPQHJciCnjubokcCEyLMmfPdvxv2nAbfgULKUlS+mNuaYwT6ztnvI7B4VuP/BDL7` +
    `Qah6EBtI4t3x16QRCisHGYmyjarnU8deMQEnOVNRvgpEhjy2vDaxs6yHdR5A0iZA/d` +
    `Yeij4bAPb2BgD0ApBIGDIjf0/SAY2kPXi4L+qD/cfm4hdHzPHUQj2/FcNAx8Dw2DFq` +
    `LAjfojb4RsJ3AR8pAfHEIUGSX1h3aXLlhZX9mtYw1TcSaVnRykuP9m34t3EpY7DLcp` +
    `S2UDlbQi9XcdW2HHRdWbne1nE1xctSTYDUFPhWD/oPt3lCmsEIv85B1+nkMH7mDoDY` +
    `YDv2mKljejwB0E4XAwusSbz/DlSgFuhPx+5CEbjdzBqB9GDfdtIETgeG4fRSgcBIGC` +
    `9fww8lDLP90w6o9QpK5UdxhGPgo7/TN0URiqcBn5vh8F/QaOaKMNu5YaMnAGtge/NM` +
    `Xr8NTA63DTRvZTF5s90GnPQRcnu8rmTkp1fZv+Yfz/H+O3Id3QRm6oy/u9GmjrJof8` +
    `hJOS4JNVTSNfDVrpqt+Zrlp1l0qqjRPB8dLrOJ9nu/NzM3xNEwhazC56RQ4r/vwjEj` +
    `W4Q/YA+uoKDmEJyL/8MaD5e4rqr8voVd0eFynLLylFVSgOIIO+uoDnMDhaKPumUB6p` +
    `B1t19GK1GxYvS+DPxOs1NX/Gy3LLW+iG9lDR8mwfUN+sjr9nFKR6/B6Vwa9kmD/lGZ` +
    `PTND0dr20h+kYI5Bkp+iekUEdQWJ/JQB0KT4uOIiX7xRJNmJTnR0PFHqrYGxqR1Pok` +
    `e8qWD0r7p60zPMs8455YzprNv4zhFNiCw094iUXCaSm3LcdqLXgSW9vv7k9CETGLJq` +
    `4c0wJmnKadTUv1ddsa3MC/55xxyIkQeEa23c7qqNJ4DWU1MZlzE8ZTwkkKpiNatTYb` +
    `CBon3rJCcpY1LbQhkUiefWTFrKZQ462/tF5GivlkwTkpZLYGnEi6JCBYMWvw3yKgu8` +
    `Q1dnNAof2k9oETQX8h1q1etWRoIcFc6obKHpY3+kONxiyP4Mkmi7wLjdrfYlGrbiR7` +
    `J+84m3FSt/v3KJb1932EB9kkGS6FyuMVOWWK92YPRI6zzLr1bnx0GEG64FhSVrQwvK` +
    `s2Nyj8G+TvoGg2yGtzixWVyRwmRK4IKUBPYoTVsndTN4rUAyUrmLDHeyJKVgi6JNYh` +
    `VW7Af1AG3lGijlJlXypAzglU4xWdfHb8DWA8WUjJCi2/ab9t7kPdozc3Y/tI1e2vOF` +
    `H5zTqr9b+LBWC8EAQeM1r8fGNmL3/atgFvx72FIHuUTTpqbxkRTjvJju4+mlboYX2I` +
    `LoBnl35D9dnl/46z1R5jTfknGoC/tAKqLuY1FdAZoWWG1zbo6QsQmbhHY/Muw2vCL4` +
    `hOc0D9/xENm1mTTvax1QH77LrWDDxB0wcl4sR0wiqZ7s3yS5eqGtFtDGVWXbPVlxJQ` +
    `c/C8KcI48Ef1nDiSIvVro9JLB+zzZ0tN8wUUccfJkrLFXsnS9JINSO0o3Wee3z04Wb` +
    `6IUn4kj3vlR0MhZkpcKaMD9tkVoRm4siLU1fHXv//rw/sbwGkKlp5GVEJLBpzkbEl0` +
    `4VVV/PuF/t51tKPcDywlx66aBgtTWqSAVZkpCYcVXiseJJvNMsUDFV0PjB1K91Wrpr` +
    `LTJ3PYdEcgZ3vZrZXuNdQ7Uyi+akqsGk6/y9hdsr0vtGivIs22N3UNHz7lBlWjescN` +
    `9O5JN9BQn4UbGCmu6QYa4yu5wU5T/gKxdk3ebfSPVWu8bXTTkO4y+s7tr+E2Zu+Q6k` +
    `raOXz7G/a71dKpmEMW7xLsgM1fTK5Tdr9IwHM94u1mlt12iWrcedonKsBXdIrNNP7q` +
    `XlEhfj232Js5X88vDm6YpS45SJECm0KjN1r/95/QfFTr5tPDbrepbrMtKVndVWDWbs` +
    `OVFHL/Wb15KmcHul2bb+8LySkRMGMwJ5zs1j5n9gQ0+6SQfA2S5GWGJYEVzTKYEPU9` +
    `ISnQQrI2V92iNhlb17KqevCerXab0QfuZwVdqftITbgTzCTLjkZpmXV2G17Ah4+8zX` +
    `+f316sgXtdqb+GDtST4QVCd9+P7jhlV3UiThmqrwrz8FHPjxL+uyCLbSVYcspeQ9GH` +
    `mwMv7GtKUf6eosyg8zPRVTV1/WzUFeypy0xRPxN1VSPdVwnjt/sarK+czfdPrQlia/` +
    `fE/LAD2Zv2JLG9fWqO2IWvNVFs7Z6YJ3Yge7c7ttv9UKP0b4KO8d85wzw9kzlR01Tz` +
    `pI6K5sAFv50xfcceG6Pkx2Zx0FXvVOw0fvazP9fb/sjHQH8yvbjGj33ML8Q6h1ytt9` +
    `BiklPZOSz7kmZhh5V/T8Qi2zq3qtfozgC/rXeuDzyhuqzO79SWnQ5lQD+7ulHwN2lq` +
    `3X77Ryb7AjNZM1m9o5wkkjW867x8lW7O/dBU45GYqQ9ABXRh3HQSbsdF/en84KiPXD` +
    `1C0qxzrg4vUZMcnZy/aAmXZtfLE1tbHUkWNdCPOCf7nqF26+BRi8uCp7sJsv31a2/C` +
    `0rVejHvmR9//CwAA//8do0KqBT4AAA==`)
  if err != nil {
  	panic(err)
  }
  container["templates/javascript.js"], err = decode(`` +
    `H4sICHKksVwA/2FtRjJZWE5qY21sd2RDNXFjdz09ALQ77XLjNpL/9RQd7u6EGsuS7U` +
    `luc/IoqWQyqc1VdjI1Tu1e1dTUBiIhCTEFMABo2Zf47z3APeI9yVU3ABKgKNlOcvPD` +
    `IwL9BaC/0GzeMA2VWsMCVo0srFAy35r1GH4ZARRKGlXxaaXWNHg5ur8cIULJl80wil` +
    `hBvmKV4WNIsGYzMBulLayUhlIVzZZLO11z+7ri+POru2/LXJRjIs9j0qJ0lDW3jZbH` +
    `cIkRctpYW5v5bGYsK67VDderSu2mhdrOfm64Qapm9tdPPvnkk0/PZxu1O7XqdM3tqd` +
    `3wU1bYhlWnmsuSa16erpS0p7sNl6fCmlOp7GnJV0Ly8lTI08IYkrcwJpaYO7EmUGtV` +
    `c23vkgXshCzVDsV/pbZ1Y3l5Ze8q3mHJpqrGOP/Wo/+DVQ3PW2JPWub5+WcXn5zPCl` +
    `YVTcUsP7X81p7uRGk3pzthN6c/sRtmCi1qO7s4Pzv/9MW/v/hT+EGLW3P7A7+1/0SU` +
    `eJVIaAK4P1fiv7hb4myGCrAWN1xOoDEcClZseAkFkzfM0NkvubVcQ831SuktkwV3eL` +
    `wyfAKF5sxykHzncUYAtMOOwGIEAIlEUz/z6680lQ/NLTqlcfS93uSZA8jG40vPB5fz` +
    `DduK6m6IVzQ7yC/GRqXIeZ4tVXmXjSeQkSataDZiWCiJG4nwJIvTCxrLs4syI0APRP` +
    `TxCPyewwlk9S1kcBLJHQhvudWiwMUH7C1nptEcxaXDI9JeKT30lBTjSQY7LSpmzHfC` +
    `2Ckry5yenBWzsiwqk5ryBIrKOF0hH5JnDirHRYgSVzTBn4h4AtnYrZ8PMjJP8ywRvt` +
    `7GcurtI8QkoKdJqflW3fBOUOS1EeWgc/NccDpi4gn7LUKxCCLr6JmN2h2hh9P79Nxa` +
    `hsiVwrBlNSihW50HKGEBWfidtehcPh47827s+Qieg9kJW2zQN+x2XEIpViuuubRwI/` +
    `jOjOD5bISxyD/DIjgbfATJtnwOy8ZaJeHbr1H/Bd+9rdhdJYydQ4b/ZRM/fsWZLjZz` +
    `yAz9aMe/FpoXVum7OWRLrXaG6yze5X8gq2hliJP49XbGjQLpY47o1yCkkzxMkfB+N2` +
    `Z+12gtbhUGnhHPGakLTngsv0/5dUcIoGDGAc3B2cIOtswWGyHXMS7hoz5ckxKEf5xW` +
    `Yt5ffzh4uh3wUnN23Q2UfMWays4jCAxLoiw5KLvh2kQzpNuP432Y5/2o+xs2w0/R/u` +
    `EG1LwQrAJGp2FgyQwvQUmwGw6GV7yw3bZ0+0i7mCWakHULK9R2y2SZe9XIJpBl40ti` +
    `x2UJLACMBqXuaDvtiwhjvkSSLxYJRHzAuFFeXdF5Z+OpW0be28wUZqWKxsQg90PC4e` +
    `C9M0Xvs1/fcBdlunRmAvau5hNYyS7J49U0AKO345LrIPPAVB4ITMBlh8j9noJ+hGQt` +
    `KzaEl2dKot9q+RJ8SHzecYPh70apUinnw91IT+qtkN9cTWDLbr+5csJRjEcHQmMhVl` +
    `oKwLyaCim5/tsPf/+ujaLsNiQ+vJoWleDSxXoE2G1ExSFfGfh84XjBs2dJxpBbypHM` +
    `GD5vSbXuwcDpAi7cwnzibOClIxQBedLuuFp5niMq6t+OA6sqtQMlOdRszd3pEjStTT` +
    `eyYJYMi7beEfKy21s7rbhc2w18Dp89RXyg/D6zSkGl5HpOaYj1WQX+c5uKDEwlCp6f` +
    `TSDidgrnHWQkotUNv4yMHLelne94x2d1lMuLMQa+ZxteVaK+zCKlHxEVg7n3tM2oFr` +
    `jjlFZ1Uc3plo6Vq1MmXsXpZZd1mK/uXmEm8IZteZ45Elmb9Qn8QfFBwALOLkHAS5TG` +
    `iX0J4uQkrNWh5rx6Lz5M4PxiAhefRdaA1CTfvRvQ/05IUvZYhf/GxXpjO7O3WqzXXO` +
    `9TcTT8laUUpsbA4iwUc3Rvq2F5ziN6YsEmOcKY1gEruQZjm9UqGHPqQAPcji+NKq65` +
    `dRcsPxlJVmzLCZTMss4lfYQBNkTjoOqopb9ACw33tHk7M0W++X9cff9maqwWci1Wdw` +
    `joOM07jHnAi/Nlp6Z+BTvN6pprsMrdebywzACToGRRieIa4fgt29YVn6Orrit2l42n` +
    `fhYWsLTyVQgyNDmBc5efdjNH1n8oCfHjIX51WJep49+Zf7Gy1BPYmcvRyB/4nhfPKs` +
    `XKbBJp2Y31KuLwUcnybGey8fQGb62kEJqzElZabV1qIEHIurGwErwqXXrr9MhxnUDQ` +
    `p0kwvZAF98CURl1mKEixYXKdYrgMkcGK7+DjdaWWrDIf4xlZzYprUrJaq4IbPKYSik` +
    `ZTxmksXkDVygFU7I5rb0XrrxtNzGAB59OzYMvr1xWrDTmvs2j0iugs/DEgoTlkNWsw` +
    `eRi5lEKaBjNXUtWJP62aM5sMGSHXVQ+KyVJtUyhywWfuQfJb6wfcIQdBtVYaA8r7D+` +
    `1YSJO/ERVPp/DUMCO74oWSZXJBoqMNGkYORuCm1Ewb/q20bh5m8G9n3sdTEs2LfZi/` +
    `dDBowQjzEs7POkfvsLIzjC6GF3FowKuDx0P+KZ6TKHsml6a+dH+RxFZIT8LlHjHyIe` +
    `wUj/4+n8XGhdAnkM0jGWnT0RduhGn3zYBuJFVB+A3XpHVrbg0UrMLEFwfvkICSJfmS` +
    `ukQVspu2QuL0DPHKoIpCUhmFfnulxcT3xnlaf5SO0lut1pqnpao2ZcszTz4bJ5G1pw` +
    `V50HZ/aDzPgigPIQbriY7bGcl75+8+UApMPykPCWb1sjO85PokZKEp1PrE3mmpWJGt` +
    `4d0H90kq9DNSmA0PuTk6X78TmElT/N/5HK9dHsxirs9RO/CI/9LeTFrAkwWcx6phuP` +
    `1BbLlqbJ5u+wSJnAW/608GL2RkloPFVK/cZg0f4a2oW7435WndmI2vsUZ2MYNvvv3P` +
    `v7+eQ6XUtYFmXd1FVlhHNywasToaCZS3rM6H7R0I42RBpE7AufnWWon6y6WG2ech2e` +
    `pUZcuEJPo9XTFWOxhXr4jAaKOSTd1TXa9KXu42mT2Lb1Bh1mzEKrk2tfufZ9k4vWQO` +
    `0l0sUsJ4qXX4Pez7CXwaHbfb6hb6gAke2p9wOK48lGxPrExO39CmmsTIuywhVg5MVU` +
    `KkgFzp9kGUY3Q/G7HeVJgmkn3hZHApLnbZSB+Y9XEIWs80aTN1spPw6KNeeHQBr32i` +
    `wBaeXORrn5Rch99BUu/JF5TReP1N426IvImIbeQdFGbeEyoE4FScEIN7IlPUHRJ13h` +
    `P6MpyGq9T9GJ9d/udfSN778Y9eq0Kpg4b7NQuXVMzTMXSjXZXBqvW64hjr3yKwp9OZ` +
    `QZzahAOMZrsMx5/m5bE6h7GqjnjjY36Al0+Y9tgk471qBe1aFxthq0oOhdUVzb3P/K` +
    `FmlEfiQeIvd3g0RgeWfTjq4aLgRDMfYpsnJ+XSlxPIvnZ1qyxaIZlpC/BapvNJ1WMf` +
    `fIBeynCPXuRl+77gS0oFrtDCD/gDb8mwEhWfANNWGDsBKyw9Vstm+68wSE89c+N5hl` +
    `uPHLLxlNCQk6h4O+/Qex7NDbYwhNgDobFgJg/D4JE5qhgsY8mjqkGe0URfGFrYCWTw` +
    `HlO4GBdHP2RJyvgYYkltzZ+I5Dvcpjeq/wZgAlxafTcBqR9/Kv3DaRNCPB+i10V46V` +
    `iG9xNKchQixzus1yKE2H9z0b4daEH2D5iGBdosvWKQOhr+ueH67opKlRhg/4Te7xWv` +
    `qi8f0oij2D8c14PjnI8f2GN0aDYDJas7V2d3muNRxAqE/dh0by9CCHiiRK3dPxHxpE` +
    `V0itw+pgrdQX3IYq/6/3hkXz/uitDLIOjRa5+/aSGX1M/5K/w7bmolTWJZavlT355c` +
    `OdnVc2ABavmTI++KRPFsm3n6yItzvcCbZ5yysAlkeAvLxvOBtLKttezFshAunaeOAm` +
    `aSCiD+1NDvIJJ3PrW/ue/nii6tqLjloKoSAmAXaJNLf5vaQhRzMcRiKLYKVmLdaA6q` +
    `IQ3fYSroJsgCai1U+9YKk0gVvQ5y7QCuorJYUCaAVzt3QUXmIAzm9jj4c8MbVzB8ub` +
    `iA//3v/wGpUlaey9Oo4wUQOSDBB6l81FHZF6dUjtgBIj6rFgbshlkolNa8sF+0IE4F` +
    `iayvNNPR0kCo+/aA/dIg0gH6j/cBoyV7aNYG/mnnq8m4OqZJAuRDkEjfO6VxCM96Ar` +
    `UyghJ8EMblF/jrDd4PerHH/UOnGiBT6sGsncIq+bHf3V4RA4QrCSc6HL/O2tNouhaj` +
    `zMkLsigURrG4M6TXKHnWRmORILudWFqJnJDCkL+rq7dU1n3s8j25h9+77kVoVpZ51h` +
    `1yynIvw4x5Ha04h7OdWnVFRfF8nFLeWxiee39ZQ7Li7WdA0lEPa39H31EucqhUnvDN` +
    `XN6STZLRgTVF84k4nRlT3Y7i/JZdc3RSGFyEhIJry4SkW4eB3N0m4AS+GA8qixbq/I` +
    `i24HR2QM+0UBfHUS+OoL44jvpiX0k/CtW4cEMa0lYUuH+0cZbYg714AuyLh2EHFZte` +
    `IXHgt3XFpMuB2VLd8AQIF5gMgL9MU2WNbqzw66+xI/UVt2fP+mjBmBG+c+cvF3AxTk` +
    `D72xdt4CPsPdrCp0C/eAz04DZ2qQ66+X3hfbpyNt8bfxpzj5E2kiQszi/+eoTJI/fj` +
    `ISYXn356hMkjj+ggk/tDJ/+wE8N/GR5Bz43ROEqNV9TjPg36fi3Zvt8pAp7O7xPhxa` +
    `EYRBwhOztEv0fw/kAgSXNjVtdcljn6wrgXJk6m+7WSvJdCRXjJi/L8eGLvunHecdNU` +
    `Nkrv2z4dP/NQ+u6AQRN0nHG6yuEvsKakZ95/l3cfSYfujxY10EoBaRk8YQeV2Arrcn` +
    `A8FKKxZbdOdDPYXNTyild5qH6esJbKUtPaSjWyzAaJ7xN+YhY7nKM6iKHk0PF6XGp4` +
    `KOQa/WVZZv1MI/Ra+sbFldIkoAFW4cX4LqS9e9c3X55MU14hi6opuXFpb999d4mN2a` +
    `imKqlrqCyBBe5WgUueQFj3qj5m/UU/Xj/sHmODRITO5Ade4DyB7l5HQ8ZK6kjoZ/v3` +
    `l6lL6NncYbcQfu+Z+jFbL0PPIKYvqbGnU8MvdiC2mw0zb5nmspdXt1W8lKJTzQN1vQ` +
    `7T1eiysnqndo56tgcSqnskRU1AezDHi4QRYN8IWqHfsC0/UPCCYOGOucvB4ItkbA7Z` +
    `dJrI/r5PfALZn8rqK9ep2S/0D15WUmFvDvXltM2fkUCp0vU7MvuHf0DtUv+WID3GwY` +
    `WbbT+lJ+ipvaupQNIJkw1twm/TrX3tonbzo5fXw5rzdN1xLrnbspTa71ONJynHfqYU` +
    `tGVvpify8bTp/uEb0G+MW78tcv3e2HU8ermtGYxhf3AUe3S8ud8LUo+IZY+mfiieRR` +
    `vRE+dQlH28txlOW+/jTgVVc/lPvryiFsyhhdJXGJLvoIXKfQugp74zUyWRTNKUehMF` +
    `Nf+NyvdvX7/pVGyvoeI+IldUKi3xD9F79d33V687gk7QpqraHpU44UQTkLxwLyf2GW` +
    `65MWx9mGXv7YNrJ6UWN4SbUvF/gCzv94AkRDsBWxoRCfqTHI/PTXzPppLU+J9Sj5IJ` +
    `6sek7wOBwUpzswFXWgZmwKgtB//1i4GcT2E9hVcbrbZ8DKYxqE6gGh0ohUaqqEGn63` +
    `YjoFat3SXrnWPtv58Iy0kbpfx6XG+Sqoe0L3zeZFUdji10kZFPxy3sxl3TJbWVNtvg` +
    `9lHRWhAiNAGqUuL/teY3QjXGw3rbTZt5alcIPiRbHVWCgwwDwnmZJhCL/xTZ3NdDPd` +
    `GQ2zHZSJq+DMe27g8UrteLEouZtLd0vTH97pUntsekrTD4+MDFPW2oAbedB3C6d/yz` +
    `GQWjEBys8g2dLjaNjn4PFn8I1UWX8NmZ/2oKWbXV2QkcbXQ51t4y0JsynNIQ2AQynn` +
    `ScDPWqHMcv45YWX3lxzbvP2+0KogIzRqwlfUIxCp22mM80y62w4eOooRAc+dB+SA2f` +
    `+u19I+X2pXuj3J8PjeV/NJ+errSLRxVhVUWf7GpVmVHr1zr3NYHOLB8wv6Pn/9AmBm` +
    `/x4ysnzBz+/Ash3v/YmkJvA8KBJ6H0vu3Fb1/uxd9njkLXZzxIn6SO3dcVXK1G/xcA` +
    `AP//Bj7F16VAAAA=`)
  if err != nil {
  	panic(err)
  }
  container["templates/style.css"], err = decode(`` +
    `H4sICHiksVwA/2MzUjViR1V1WTNOegC0GXtv27j9f38KTocbmiB2HtfsGgUDtnUdMG` +
    `C9Fc0wDBjuD1r6yeZCkRpJ+XGHfPeBL0mkRdlOOhR1TP7eb5K+vkR/+eu/Pn/KEeX8` +
    `uW1QK6FqKfpDDSXB6L8tCAISXV7PZkte7tGvM4R6miUoBQJVnCkkgUKhCGcG2WAZwL` +
    `zCNaH7HGVf+ZIrjj5zxrMrVHPGZYMLeLy8nh2gktVaAW9ldoWKVkiygceZx5LkF8jR` +
    `3Ydm9zhDqMZiRViO3ttlg8uSsFWObuy64JSLHG2weDefV1zASvCWlRcatsTFs116hH` +
    `7n4nH2Mps5N7yrCZtvSanWOfrxdx+a3YVxROcRbe2Q2XdQ6X/WslGtXzR3wppW/Vvt` +
    `G/h9pmCnsp+vgj3W1ksQ8W6DpdxyUWY/G9lOrdubm++NTVyUIHJ02+yQ5JSUCdMN2l` +
    `zgkrSyc96S77Sexn8OY8kNZL6F5TNR8wmMmv8yAR744BZq49zrS1TwuuYMSbWnNssW` +
    `S7IyZg3x77RpL7PZYslp2UO3oLNEC6KlhetotfUB/Y2nlzWmNAbfLH68dxotSMGZhR` +
    `NKc8Q4g8eBjzUaQmsn1y35BkRF+TZHa1KWwPReSWRD8T5HS8qLZ70jleDPkEhEC/QZ` +
    `9uH7wR4lDP7DdYIbdO8H7VoofTniskTYIji3o98iCQqpNQyyHJl89Bly95oM8dV2G1` +
    `TbfMmV4rXDGq2s4v7C6777CrLhTBe1N0BuiSrWhK10S9kCMF23bc3Qu5ovCYULhFmJ` +
    `BN+idyXIZ8WbC2tM5+iKgs0zCrt5SYTtRbnjoyGYkhWbEwW1zFEBTIGwGq1JaRVpuC` +
    `SWCi8lp60C9BtSN1wozJRmoXiTo/nDw8NDs4tAFCo1DnuZzZatUi6xuv50Fzr07rAC` +
    `WyZBPb41YpO5q7vO3Dimd0lUHA8WcyKijSDc9cOTeo81eS6sJkbNzkf5WpeT4cYbXB` +
    `C11zrcn9atHYuKF618G4uSSLykrsBiLsFIKd4nRklxd8oIGSmHwxwWfNvNDE3w0daG` +
    `L366xXvpC+a8qvAFybcRN11qp7FyyoV5PeC7qLANxgaEIgWmPt8UbyxiAa4pR7LsiK` +
    `gxYWgliOleC7MKcDVIK6D/zhXUDcUK5tZAWwXothK+GkIsLADL3Pg8WxhJn4TgAi2y` +
    `4d5HzpTgdGQXmEKL7JQwd3qfqujwmEXxnrcKudNE0oRxI1JmjBrik6zn8qt3mhaU9+` +
    `zjpP+uqm6r2+qxI/fiRhg4UIirVUjggmuj15cIWIl4FSaFzxIvMUiWoRqn5gxuFbcf` +
    `OiDdajp/CiXoPwlskf7yxNlq4kv2bUi+ULwHgRZm8ZmX0H3JXGV57Mixfr/HMrJGsP` +
    `R+j/WJ4ka6rhghOtBg4ry3gyMYJvd2z45iCbTKdURTk8iL/XMrsLlZjMj1sEPBJ8y2` +
    `lDpesHPwiFgL6TGN90fw9P5I7sapas7CPlhd/gbhOyd5T8vUvxGporR6AiyKdbT5J8` +
    `G3ErLh8eEHf2Ae8kplmYaF2FZMCt9CQwqrQ4rCQiM/H/jTO9lFNXDzINLnd4mjHcL6` +
    `0srQnwM1nhRvBssvAjaEt3Kw9ZO+GT6a6Xw0rEkR2QHOhCCk3RLWgGGZrAP9GVeNMS` +
    `xJ8dRP/0OV0nIcRkypNU9TaehIbkRp4LPDVHKQG11t/98yQ0v4yJls676Ff4UGsOqW` +
    `T4St6ACKWcnrLGxAnkWiDzlwSOPEJEgsNKRwmiQoLDSSYZRNyjBQS6H4akXh77bTmw` +
    `ZeQsFth9cXohKEvhCPRDMIm4+lmWorQbpQdmOuO66645ZvbO/9TDgr1v6YMN13caHI` +
    `BrQK/yCKAjqyzmKqPwqlu+zRjUNCumxrdGQ96BB+zPcYXwRfCZASBTPXpV9sRxTnCJ` +
    `wezcMxPDwCHBicFGDhb5JgXJMWoMFv4d85MinCY5zxmDcudkxqSTb2pj59Ee7r4Wbh` +
    `TlP+ccE9H0wp9KLLDyOpa40whBHVQdMlqLf8XfNVh/F0G9Ws/alKGpMJN18+AqXZ4N` +
    `kheKo6fnEbqny6jq/XclxPfy/z2gwelHyOuPfuYWTt8+VLT7boZI69c5yUbbbxYmRf` +
    `ALjYRyHu9l8V58kQd6y9B/uNt8X4QOeTlHyDmtNBDtQ5M9IB7TcJ9+LAnqh1xfCIyD` +
    `/ujFIApXGHvJotGOxUN6mvL8ce124vhm+s/rzqKykW6PfPGezDme6fuc37792xlyQv` +
    `zU5Dv+qGZojkRlq4jHHMUApWpzw6XV8i91hun45C90wa7v97B0z2s85UlDIq0H9weh` +
    `i2tZ5NInjd6SE1ZSPpCTb9GeEon7GzQAA8hUvq8SKGx6cKM2ljAVJhoYYZ6eaxf3S9` +
    `vkQt24AgFanAHoUXujWTw9cdt22VdYtv37MdY9863PITU2KfBaITzSWEnlm/02Uasv` +
    `4iSI3FPhuDPUHBWYm9ykemyYhBx2qtr4ShgxK+MLDXdrK0OzQrLcHWoF9FXctvuxIL` +
    `lzGOKZ9g5TC6/v8TrgEFqyx6o/vgfro9s9MduDHpnB8qge4qYf4e73j+5XzcWbFXDh` +
    `2Q8Ou0T05fDVpqqGGUSQHwMcAfb54hNKIYa5MBMMRPNcQYbqhKshk0p5yp9bxYE1q+` +
    `gw2wiysDd6fbGJi46PifLac5xyemaUljx6vDXwmnhPKyTFqjYSljTuF7qi1juCOm3H` +
    `e3AOCVrrf/BQAA//+WKnubRCQAAA==`)
  if err != nil {
  	panic(err)
  }
  container["templates/theme-default.css"], err = decode(`` +
    `H4sICNOasVwA/2RHaGxiV1V0WkdWbVlYVnNkQzVqYzNNPQBcikEOgjAURPc9xSTdkZ` +
    `SqqGnKacpvESL4SVvThfHuBiQunLd7b2xkzngJQFcgXsbg0UeeMeS8JKt1KaUuTaKB` +
    `eUo18aynsdOlUXkIc1Ax+JpSQqWFAJTqHN1vkZ8PbyFDv9JuoecYfuGw7RvoaCGNN2` +
    `RoFycLeTUruzj/Py4W0jlnnGnF+xMAAP//Q68dl8MAAAA=`)
  if err != nil {
  	panic(err)
  }
  container["templates/theme-juri.css"], err = decode(`` +
    `H4sICNOasVwA/2RHaGxiV1V0YW5WeWFTNWpjM009ADzMQQ6DIBCF4T2nmMSdCaJVUf` +
    `E0OkA11Y4BGhZN795oSWf5f5OnHFGANwMQOSAdq9FgHe2whHB4JUSMsYi1x4Vo8wXS` +
    `LrZ1FrHmYTG74c7oAr2HXDAGwPk84ePu6PXUCrJOo+7b8QJLzvyhvO4HWCnIZIlGzi` +
    `ncFGT9oJuhSqE5P5oB5ZRCe45LtJ0d2ecbAAD//9Kzz8TDAAAA`)
  if err != nil {
  	panic(err)
  }
}

func Content(filename string) (*[]byte, error) {
  if file, ok := container[filename]; ok {
    return file.content, nil
  }
  return nil, fmt.Errorf("file2go %s not found", filename)
}

func ContentMust(filename string) *[]byte {
	content, err := Content(filename)
	if err != nil {
		panic(err)
	}
	return content
}

func ModTime(filename string) (*time.Time, error) {
  if file, ok := container[filename]; ok {
    return file.modTime, nil
  }
  return nil, fmt.Errorf("file2go %s not found", filename)
}

func ModTimeMust(filename string) *time.Time {
	modTime, err := ModTime(filename)
	if err != nil {
		panic(err)
	}
	return modTime
}

// eof