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
    `H4sICNOasVwA/2FXNWtaWGd1YUhSdGJBPT0A7Fvdj+O2EX/PXzFRXxpEkkXJsq2F5e` +
    `Jyd20K5NLD3uH68UZLtM2EEl2SttcJ8r8XFGV92PLXxrt7VwQoukdqOJz5zXA4nInH` +
    `X7/5x+uP/37/FhYqY5OvxvoPMJzPY4vk1uQrgPGC4FT/A2CcEYUhWWAhiYqtlZo5I6` +
    `v5KccZia2MTykjzoZMHbxcOgle4ikjFiQ8VyRXsbUlsmPZmpLNkgvVINzQVC3ilKxp` +
    `QpxiYAPNqaKYOTLBjMTILvgAyIWg+c+O4s6Mqjjnuw0YzX8GQVhsSbVlRC4IURYsBJ` +
    `mVM24i5SXEakEy4qRkhldMNRd97TiN1YZ4odRS3vV6M54r6c45nzOCl1S6Cc96iZR/` +
    `meGMsm18z6dc8W/f8ZxbpSL720/KeceZPH4bOl8owlfy9CZjRRUjk3+SKbx7/3rcM0` +
    `PtA72dE4ynPN0CTWNL/6MBASyxlLDEAmdEESFhJngGGaa5O+egOCxompIcaL5cKZhR` +
    `wlJZ65TSNSQMSxlbC5qSSumxIdf7baQFarsksaXIg7JgjdmKxNavv7ob+dtvO0l6KV` +
    `1Pvqql+vDpb0ATnpu9zLxczwELih0jUmwpsSIWzHiyktpTY2uGmSQWFADFVkrlkuHt` +
    `Xc7zhmRym005K0TTGzgJFYn2cu3G3/GH2PLAA1//r1oDMDZUkDzEFvIsSLbmr4itkT` +
    `UZ98znao+e2eTonlqsczsusVpAGlvvILAD+ARoCD8AGtrIg//oLfX3yzfEK3lWR0ES` +
    `BeboWn0LFkR7X2yhvgUPxcw2tgK9tSa8eB3yuxeelTk37nIhSH4NUtSJUYsaoQamo0` +
    `diKsj6cvnQqLEjOi9hVJP7j5NPKr68xuSj2nSjwnKDwnCDqw0nCRbJ4nJotL3cEAIb` +
    `DTsgaZ68wJy8YXHwQjd8xNFjVF7hVdpNvoc+fIIAgpP2GhSUaKRJT7ueHZUsI4jOsI` +
    `x2LE8TIhuFJU8UAgrPcC2IC7aHpGfxm3GWEnEpghn4tg8LCIFBYPtuCAsoNh7B9+A/` +
    `5tDxzfnd56AEzuWMi0xfEjiXDCvyZ892/G8adDs5JU/pkqf0zlxTWCRWPaf9zkGh2w` +
    `/80AutxmJoEG1jy3eHXhCEKGwsFuYU1ePiPHXzkAoLHTUb5CRPY8tr0xc2dJDvosgb` +
    `RMgeusPQR8NhH17BwB6ABgSBgyI39P0gGNpD14uC/qg/rD+3GDq+5w6ike14LhoGvo` +
    `eGQYtR4Eb9kTdCthO4CHnID44xigxI/aHdhQVfVld2a1nDVIIrbScHaem/OfTivYDl` +
    `DsM6ZOlooINWpP9uYyvsuKh684v9bIrzm6YE+0fQ00ewf9T9O9IUnstVdvYOv8yhA3` +
    `cw9AbDgd80RcubUeAOgnA4GF3jzRf4cgmAGyG/H3nIRiN3MOqHUcN9GwwROJ7bRxEK` +
    `B0GgaT0/jDzU8k83jPojFOkr1R2GkY/CTv8MXRSG+riMfN+Pgn6DR7RDw660BgbOwP` +
    `bgl6Z6HZ4aeB1u2oh++mKzB0XYc9DVwa60uZPSIr9N/zD+/4/x25RuaCM3LNL7gxyo` +
    `dpNjfiLIkuCzWU0jXg1a4arfGa5aeZcOqo0VwenU67ScF7vzUwt8SxNIms+vekUOS/` +
    `n8Exo1pEP2APr6Cg5hDci//jFQyPcY6G8r6E3dHucpz65JRfVRHACDvr6AFzA4mSj7` +
    `JlEe6QdbufRq2I2I1wXwJ5L1lshf8LKsZQvd0B7qvTzbB9Q3o9PvGU2pH78ndfBLHR` +
    `aPecZkNE3Pn9e2En2jBPKMFv0zWuglKKzWMNCLwvOqo0jrfrVGU67U5aehFA+V4g2N` +
    `Snp8Vjxty08a/fPWGV5knnFPrufN4h/jOAW+EvATXmOZCLpUdcmxHEuRxFb93f1J6k` +
    `3MoMkrwzSHuaBpZ9FSf61Lgzv6t0JwARmREs9JXe0sl2rEKyqrycmsm3KREkFSMBXR` +
    `srTZYNBY8ZrnSnDWtNBui0QJ9oHn82qHim/1pfUy0sInKyFIrtgWcKLomoDk+bwhf2` +
    `uDokpccTcLNNuPeh4EkfQXYk2KUUuHFhMsVFFQOeDyqvhQsTHDE3zYdJV1sdHzNRc9` +
    `Os6EMLyUOsqWbDRQb80cyAwzZk28Ox8dZ5CuBFaU5y0Ob8rJHQv/Dvl7LJrl68oYck` +
    `NVsoApURtCcij6JNJqWaPpDnqrT5RsYMof7olc8lzSNbEO5Nwj/0HDP2m9qM0Z0uhT` +
    `CWpBoGx+FKFhzxsAxtOVUjwv9DfFsd1tVVTQzb3VXlLW4ktJdPSxLirM73MBGK8kgQ` +
    `dG85/vTGfkT3WRbjLurSQ52NkEi/aUUaEFVaeN97D7YAqVx/GQXQRPrv1u1yfX/zvB` +
    `NweCNfWfFgTiuQEoa4y3BKDzhC4Z3tpQ9EaAqMQ9eTbfM7wl4orTaRbo/z+BsOkEFa` +
    `E4tjponxzrQoBHIH1UI0FMnarU6d4Mv3StygbazlBm1NX5fC4FCwmeNkQYB/6gk/0T` +
    `IbJ4C5S4dNA+fbQs9nwGIN4LsqZ8JU+d5h1J5Sjda57ePQRZPwsoP5KHg/SjAYjp4Z` +
    `ZgdNA+ORCFAE96kezB8o6n5NQloe+dv/79X+/e3sGM5ilgnSAqImCDt6A4KD6fMwJq` +
    `QWVX4r63031ZAikR/mgWm6oDZPwgLrUCdUH1xqR4LxrMykLO7zJTl25v80K1F9Gmrv` +
    `ncwvvOuUFZAN5zg2L2rBsUVJ+FGxgtbukGBccXcoO9YvcVau2bvNvoH8qSc9voptDb` +
    `ZfS9e7ug25m9Q6sboXP83jbid8PSCcwxi3cpdsTmz6bXObtfpeClHvF61yNuu0TZRj` +
    `zvEyXhCzrFrst9c68oGb+cWxz0cm/nF0cnzLBIOUieAp9Bo+ZY/XeV0HwOF2WjT/t1` +
    `oqpAtqZk874ks/YLmSRXhw/i3SOXHalT7b69zZWgRMKcw4IIsp/7XPiaL8QnuRJbUC` +
    `RbMqwIbChjMCX6e0JSoLnibam6VW0Ktq10lTyf3/PNfpH3yP2sqUu4T+SEe4eZMHby` +
    `lC5ZZ53gGXz4xKv69/nt1Qjck4wfoPgsGExx/hxH99CP3gvKb+pEgnJUXRUaTvP8WM` +
    `J/V2RVZ4JLQflLAH38Wf/MvqaB8g+AMg3EzwSrspv52cAVHMBlupOfCVxlq/RFjvHr` +
    `QwSrK2f3/WOrM9eaPdOX62D2qt2ha0+f68918Wt16lqzZ/p0Hcze7Dfc9j9ULP27oK` +
    `Nxd0kbruimnMlpyk5QR0Zz5IKvu0Pf8YdGi/ahmRx05TulOI2f0xx25Oofzxjqj6aK` +
    `1vgRjfnlVWd7qvUWWk0zqjrbXF9SF+s4+PdErljt3Dpfo3uN8TbuoljwiOyyXL+XW3` +
    `Y6lCH97PJGKV6lqTX59o9I9gVGsmawekMFSRRveNdl8SrdrfuhCeOJM1MtgJLoynPT` +
    `uXH7XFSfLj8c1ZKbn5CUdXbE4TlykpM972dN4VJ2uzhR2+pEsKiIfsQZOfQMPVsdHj` +
    `247vB0F0HqX5X2pjzdFoNxz/yY+n8BAAD//4/4kFldPQAA`)
  if err != nil {
  	panic(err)
  }
  container["templates/javascript.js"], err = decode(`` +
    `H4sICNOasVwA/2FtRjJZWE5qY21sd2RDNXFjdz09ALQ77XLjNpL/9RQd7u6EGsuS7U` +
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
    `4d0H90kq9DNSmA0PuXlL+GQB5/1Uo1XPdp3wF7hAySLtQPdNhtyvSWTLSsjrbJyo2F` +
    `EsXyPoITqT5vYHseWqsXl6fhPU1bPgwP0R482O7HuwKuutxKzhI7xedUvxPmFaN2bj` +
    `i7WRgc3gm2//8++v51ApdW2gWVd3kTnX0VWNRqyORgLlLavzYccBhHGyIFIn4OJFa/` +
    `ZE/eVSw+zzkLV1OrdlQhL9ntIZqx2MK3xEYLRRyabu2YDXSS93mxWfxVexMGs2YpXc` +
    `v9r9z7NsnN5WB+mmGuVuxw6/h30/gU+j43Zb3UIfsOVD+xMOx9WZku2JlcnpGxpnk3` +
    `iLLt2IlQNznhByIFe6fRDlGP3YRqw3FeabZKg4GXyTC4I20gdmfUCD1sVN2pSfDDI8` +
    `+vAZHl3kbJ8oQoYnF0LbJyXX4XeQ1NvrglIjr79pAA8hPBGxDeGDwsx7QoVInooTgn` +
    `lPZArfQ6LOe0JfhtNwJb8f47PL//wLyXs//tFrVaiZ0HC/+OHc0zwdQ3/clSusWq8r` +
    `jknDWwT2dDoziHOkcIDRbJcq+dO8PFYwMVbVEW98zA/w8pnXHptkvFf2oF3rgixsVc` +
    `mhsLqiufeZP9SMElI8SPzlDo/G6MCyD0c9XBTlaOZDbPPkpFwedALZ164AlkUrJDNt` +
    `AV7LdL4XX/rgA/RShnv0Ii/b9wVfUk5xhRZ+wB94S4aVqPgEmLbC2AlYYemxWjbbf4` +
    `VBeuqZG88z3HrkkI2nhIacRMXbeYfe82husIUhxB4IjQUzeRgGj8xRxWAZS55kADTR` +
    `F4YWdgIZvMdcMMbF0Q/ZocTgILGkSOdPRPIdbtMb1X+VMAEurb6bgNSPP5X+4bSZJZ` +
    `4P0esivHQsw4sOJTkKkeNl2GsRQgykN+E1Qwuyf8A0LNBm6V2F1NHwzw3Xd1dU88QA` +
    `+yf0fq94VX35kEYcxf7huB4c53z8wB6jQ7MZKFnduYK90xyPIlYg7Memew0SQsATJW` +
    `rt/omIJy2iU+T2MVXoDupDFnvV/8cj+/pxd41eBkGPXvv8lQ25pH7O1wLecVMraRLL` +
    `Usuf+vbk6tKuMAQLUMufHHlXbYpn28zTR16c6wXePOOUhU0gw+tcNp4PpJVt0WYvlo` +
    `Vw6Tx1FDCTVADxp4Z+B5G886l9CWA/V3RpRcUtB1WVEAC7QJtUD9rUFqKYiyEWQ7FV` +
    `sBLrRnNQDWn4DlNBN0EWUGuh2tdfmESq6L2S6ytwpZnFgjIBvCO6my4yB2Ewt8fBnx` +
    `veuMrjy8UF/O9//w9IlbLyXJ5GHW+SyAEJPkjlo47KvjilcsQOEPFZtTBgN8xCobTm` +
    `hf2iBXEqSGR9yZqOlgZCAbkH7JcGkQ7Qf7wPGC3ZQ7M28E87X03G1TFNEiAfgkT6Ai` +
    `uNQ3jWE6iVEZTggzAuv8Bfb/B+0Is97h861QCZUg9m7RRWyY/97vaqISBcbTnR4fi9` +
    `2J5G07UYZU7etEWhMIrFnSG9RsmzNhqLBNntxNJK5IQUhvxdXb2l+vBjl+/JPfwCdy` +
    `9CU9miO+SU5V6GGfM6WroOZzu16oqq6/k4pby3MDz3/rKGZMXbz4Ckox7W/o6+o1zk` +
    `UM094Zu5vCWbJKMDa4rmE3E6M6YCIMX5Lbvm6KQwuAgJBdeWCUm3DgO5u03ACXwxHl` +
    `QWLdT5EW3B6eyAnmmhLo6jXhxBfXEc9cW+kn4UynrhhjSkrShw/2jjLLEHe/EE2BcP` +
    `ww4qNr2L4sBv64pJlwOzpbrhCRAuMBkAf5mmyhrdWOHXX2NH6ituz5710YIxI3znzl` +
    `8u4GKcgPa3L9rAR9h7tIVPgX7xGOjBbexSHXTz+8L7dOVsvjf+NOYeI+1ISVicX/z1` +
    `CJNH7sdDTC4+/fQIk0ce0UEm94dO/mEnhv8yPIKeG6NxlBqvqMd9GvT9WrJ9v1MEPJ` +
    `3fJ8KLQzGIOEJ2doh+j+D9gUCS5sasrrksc/SFcVNNnEz3ayV5L4WK8JI37vnxxN61` +
    `9bzjpqlslN63DT9+5qH03QGDJug443SVw19gTUnPvP9S8D6SDt0fLWqgJwPSMnjCDi` +
    `qxFdbl4HgoRGPLbp3oZrBLqeUVr/JQ/TxhLZWl7reVamSZDRLfJ/zELHY4R3UQQ8mh` +
    `4/W41PBQyDX6y7LM+plGaNr0HZArpUlAA6zCi/FdSHv3rm++PJmmvEIWVVNy49Levv` +
    `vuEhuzUU1VUvtRWQIL3K0ClzyBsO6df8z6i368ftg9xgaJCJ3JD7zAeQLdvdaIjJXU` +
    `2tDP9u8vU5fQs7nDbiH83jP1Y7ZehuZDTF9SY0+nhl/sQGw3G2beMs1lL69uq3gpRa` +
    `eaB+p6Haar0WVl9U7tHPVsDyRU90iKmoD2YI4XCSPAvhG0Qr9hW36g4AXBwh1zl4PB` +
    `F8nYHLLpNJH9fZ/4BLI/ldVXruWzX+gfvKykwt4cavBpu0gjgVKl67d29g//gNql/i` +
    `1BeoyDCzfbfkpP0FN7V1OBpBMmG9qE36Zb+9pFfetHL6+HNefpuuNccrdlKbXfpxpP` +
    `Uo79TCloy95MT+TjadP9wzeg3xi3flvk+r2x63j0clszGMP+4Cj26HhzvxekHhHLHk` +
    `39UDyLNqInzqEo+3hvM5y23sedCqrm8p98eUW9nEMLpc85JN9BC5X7XkJPfWemSiKZ` +
    `pLv1Jgpq/mOX79++ftOp2F5DxX1ErqhUWuIfovfqu++vXncEnaBNVbU9KnHCiSYgee` +
    `FeTuwz3HJj2Powy97bB9eXSs1ICDel4v8AWd7vAUmIdgK2NCIS9Cc5Hp+b+OZPJekL` +
    `gpR6lExQYyd9aAgMVpqbDbjSMjADRm05+M9oDOR8CuspvNpoteVjMI1BdQLV6ECpDr` +
    `1pXYNO1zZHQK1au0vWO8faf4gRlpM2Svn1uN4kVQ9pX/hOyqo6HFtoRyOfjlvYjbvu` +
    `TepPbbbB7aOitSBEaAJUpcT/a81vhGqMh/W2mzbz1K4QfEi2OqoEBxkGhPMyTSAW/y` +
    `myuc+QeqIht2OyuYa2ngzHtu4PFK7XixKLmbS3dL0x/e6VJ7bHpK0w+PjAxT1tqAG3` +
    `nQdwunf8sxkFoxAcrPKdoS42jY5+WBZ/UdVFl/D9mv/8Clm11dkJHG10OdbeMtCbMp` +
    `zSENgEMp50nAz1qhzHL+OWFl95cV3Az9vtCqICM0asJX2LMQotu5jPNMutsOErq6EQ` +
    `HPnQfkgN3wzufWzl9qV7o9yfDx3qfzSfnq60i0cVYVVF3/5qVZlR69c69zWBziwfML` +
    `+j5//QJgZv8eMrJ8wc/vwLId7/2JpCbwPCgSeh9L5t6m9f7sUfeo5C12c8SN+2jt1n` +
    `GlytRv8XAAD//33AwPTuQAAA`)
  if err != nil {
  	panic(err)
  }
  container["templates/style.css"], err = decode(`` +
    `H4sICNOasVwA/2MzUjViR1V1WTNOegC0Gelu47j5v5+C1WKLJIido5PuREaBttMpUK` +
    `CzDSZFUaDYH7T0yWZDkSpJ+dhF3r3gJYu0KNvJLAI4Jr/7JumbK/TXv/37y+ccUc5f` +
    `2ga1EqqWoj/WUBKM/teCICDR1c1ksuDlDv0yQWhPswClQKCKM4UkUCgU4cwgGywDmF` +
    `a4JnSXo+wrX3DF0RfOeHaNas64bHAB86ubyQEqWa4U8FZm16hohSRrmE88liQ/Q47u` +
    `Pzbb+QShGoslYTn6YJcNLkvCljm6teuCUy5ytMbiYjqtuICl4C0rLzVsgYsXu/QI+5` +
    `3L+eR1MnFuuKgJm25IqVY5+uH3H5vtpXFE5xFtbZ/Zd1DpP2vZoNavmjthTav+o3YN` +
    `/CFTsFXZT9fBHmvrBYh4t8FSbrgos5+MbKfW3e3t98YmLkoQObprtkhySsqE6QZtKn` +
    `BJWtk5b8G3Wk/jP4ex4AYy3cDihajpCEbNfx4B93xwB7Vx7s0VKnhdc4ak2lGbZbMF` +
    `WRqz+vj32rTXyWS24LTcQzegs0QLoqWF62i19QH9raeXNaY0Bt/OfnhwGs1IwZmFE0` +
    `pzxDiDec/HGg2hlZPrlnwNoqJ8k6MVKUtgeq8ksqF4l6MF5cWL3pFK8BdIJKIF+gz7` +
    `+H1vjxIG/+U6wQ268wMl7MWm4EEKF/eX3lna/1D6msVlibDl4mKDfoskKKRW0OODTN` +
    `L6NLp/Sxr5krwLSnK64Erx2mEN6/7Q6b79CrLhTFe+N0BuiCpWhC1139kAMF3cbc3Q` +
    `Rc0XhMIlwqxEgm/QRQnyRfHm0hrTRaOiYJORwnZaEmEbVu74aAimZMmmREEtc1QAUy` +
    `CsRitSWkUaLomlwgvJaasA/YbUDRcKM6VZKN7kaPr4+PjYbCMQhUoNw14nk0WrlMu+` +
    `rondhw69PyzTlklQ8/dGbDTBdWuaGsfsXRJV0KPFHIloIwh3TfOkBmVNngqriVGz81` +
    `G+0jVnuPEGF0TttA4Pp7V0x6LiRSvfx6IkEi+oK7CYSzB3ig+JeeOLdXzODJTDYQ4L` +
    `vukGiyb4ZGvDFz/d4J30BXNeVfiC5JuImy6101g55cK87vGdVdgGYw1CkQJTn2+KNx` +
    `axANe5I1l2jtSYMLQUxHSvmVkFuBqkFdD/pwrqhmIFU2ugrQJ0VwlfDSEWFoBlbnye` +
    `zYykz0JwgWZZf+8TZ0pwOrALTKFZdkqYO71PVbR/FqN4x1uF3JEjacKwESkzBg3xSb` +
    `bn8ot3mhaU79nHSf9dVd1Vd9W8I/fiBhg4UIirVUjggmujN1cIWIl4FSaFzxIvMUiW` +
    `vhqn5gxuFbcfOiDdajx/CiXovwhskP7yzNly5Ev2bUieKN6BQDOz+MJL6L5krrI8du` +
    `RYv7/HMrIGsPT+HuszxY10XTFCdKCD8fFgx4cdvhJolesY7ln+pRXYXC0GeHpYPKn0` +
    `rD1bkHPWgBgL2WMaTw7g6f2BPIzTzhx+veO7XAxCcU4inpZ1fydSRSnyDFgUq2jzz4` +
    `JvJGT9o8AHf0Lu80pljIaF2FZMCt9CQwqrQ4rCQiM/H/jTO9lFNXBzL9LnV/zRare+` +
    `tDL0Z0+NZ8Wb3vJJwJrwVva2ftRXwbmZtEfDmhSRHeCMCELaLWENGJbJOtCfcdUYw5` +
    `IUz/tJfqhSWo7DiCm15mkqDR3IjSgNfHaYSg5yo6vtXy0ztIRPnMm23rfjr9AAVt3y` +
    `mbAl7UExK3mdhQ3Is0j0IQcOaZyYBImFhhROkwSFhUYyjLJJGQZqKRRfLin8w3Z207` +
    `5LKLjt6PpyU4LQN+CBaAZh87E0E2opSBfKbmR1R093dAoa25mR1gPfD/3xzosLRdag` +
    `lfgnURTQmess5vInoXTfPX+jO5X2AHTR1gfrWa9x+El+jKY/hl2GxoZFqRCB07fM/q` +
    `TuX8cPDE4KsPB3STBGpgVo8Nn89S0KSZ2dhCGMqLZBJ63e8jetNx1F041Hs/bnEP39` +
    `SV/JzSZQmvUu3cFDzfFrS1/l03V8u5bDevpbidem95zinxzck3D/RmJf+F73ZLNO5t` +
    `At/6THCx9de//lYheFuNt/U5xHQ9yx9h7cb7wvxgc6n6TkO9QcD3KgzpmRDmi/Sbhn` +
    `B/ZEzSKGR0T+aWOQAiiNG9L1ZMZgq7rZdnM19LR0d9l/YfQnPF9JsUC/f84w7M9A/8` +
    `hrXj/vj72jeGl2OPhVN0NCJNfhw2WMY3p0sDrlyeXmCrmnYvtwErpn1PD4EDDazzpT` +
    `UcqoQP/eMO23tT2bRPC6YZoaapH0BJv9yDzKZ2g0BsBTuKSu9zE8HrLmcTgWIBUWqp` +
    `+R7gnZPzneXKGWrUGQilRgD48z3ZrJ4duG27bKusW379mOsW8dbvmZKbHLAtGJ5hJC` +
    `z6zf8TINWT8JUmOxy4Zgz1BwVmKv8pFpMmDQsVrbV0LfQQlfGNhbO1naHZqVlmBr0K` +
    `+iruW3XYmFyxjHlE+wchhd//8R14CCVRa9an10v26e2ekO3Jh0zu8qge4rYf4f73j+` +
    `3XjYWbFXDh2Q8Ou4T05f9VpqqGGUSQFwHuAPN88QGlEMtckAGOKnGmIMN1QlWfeaU8` +
    `7UalqsCC0vYA3s8trA3ek2BqZ+QnY/2o1zjk9M45KGjleHv5GNCeVlmbRGw1LGnML3` +
    `VFuGcAdMeehuAcArXW//DwAA//8W5vpLZyMAAA==`)
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