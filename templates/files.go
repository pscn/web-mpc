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
    `H4sICDM4slwA/2FXNWtaWGd1YUhSdGJBPT0A7Fzdk+I2tn/PX3Hi+3JTsdWWv90FfW` +
    `syM3ezVZlsV8/U7MebwQKUGIuVBDRJ5X/fkmSMDQZDh+6ebKVqakDy0fn8STrSMT34` +
    `+t3f3n765/17mMl5cffVQH1AkZXToUVK6+4rgMGMZLn6AjCYE5nBeJZxQeTQWsqJk1` +
    `jNR2U2J0Nrzka0IM6ajJxssXDG2SIbFcSCMSslKeXQ2hDRMWxFyXrBuGwQrmkuZ8Oc` +
    `rOiYOLphAy2ppFnhiHFWkCG2NR8AMeO0/NmRzJlQOSzZVkBBy5+Bk2JoCbkpiJgRIi` +
    `2YcTKpetBYiHOI5YzMiZOTSbYsZHPQ147TGG2IZ1IuxO3NzYSVUqApY9OCZAsq0JjN` +
    `b8ZC/N8km9NiM3xgIybZtx9YyazKkH3xd1W/49w9XQydziRhS3FayEBSWZC7v5MRfL` +
    `h/O7gxTYWBmy0IBiOWb4DmQ0t9abgAFpkQsMh4NieScAETzuYwz2iJpgwkgxnNc1IC` +
    `LRdLCRNKilzsbMrpCsZFJsTQmtGc1EYPDLmStxYWyM2CDC1JHqUFq6xYkqH1669oLX` +
    `77bavJTU5Xd1/ttPr4+S9Ax6w0sky/WE0h4zRzjEpDS/IlsWDCxkuhkDq0JlkhiAXa` +
    `QUMrp2JRZJvbkpUNzcRmPmKFVk0JcMaUjxXKFYy/Y49DywUXPPWvHgMwMFQwfhxa2L` +
    `VgvDGffGgl1t3gxjyuZdwYIUdlKrX6JC4yOYN8aH0A3/bhM+AYfgAc29iFfymR6vn5` +
    `ArOl6LWRk7EEM3WtwIIZUegbWjiw4FH3bIaWr0QrwrPHYa97YK/OpYHLmU7ydk5KO3` +
    `3Uosa44dPkiT7lZHW+fjhpSMT9GqY7cu9p+gnJFpeEPNmFLtGRi3TgoosDJ0jGx7Pz` +
    `XaPihULwbRx3uKQ583wz82I98UIUPmHqFVRcgCoFk+8hgM/gg38yXpGmxIkiPQ09O6` +
    `1YppD2sEy3LE8TYhuHFU8cAg57uGpizfaQtNd/E1bkhJ/rwTl4tgczCKEA3/ZQCDPQ` +
    `ghP4HrynTDq27pc+BcmzUkwYn6tNIitFkUnyv67teN806LZ6CpbTBcvprdmmMj62dn` +
    `0Kdw4OUeB7oRtajcHQINoMLQ/Fru+HOGwM5mYW7dp6PnXzEDLjatVskJMyH1pum17H` +
    `0MEewqkbpdiOURx6OI4DeAORHYFyCAYHpyj0PN+P7Ri5qR8kQbx73GLoeC6K0sR2XI` +
    `Rj33Nx7LcY+SgNEjfBtuMjjF3s+ccYpcZJQWx3+YIt6i27NawRKs6kipODlfbfHKJ4` +
    `b8FCcbhbstRqoBatVH1uhlbYsVHdTM/G2Sgrr5oS7E9BV03B4Cj8O9IUVorlvHcPPw` +
    `/QPopiN4ojrxmKFpqxjyI/jKPkEjSfgeXKASjFXpC62MYJipIgTBvwbTDE4LgowCkO` +
    `I99XtK4Xpi5u4ROFaZDgVG2pKA5TD4ed+AwRDkM1XRLP81I/aPBIt96wa6uhACeyXf` +
    `ilaV4HUn23A6aN1U9tbHaklz0HX7zYVTF3cqrz2/zP4P/3BL9NiUIbo1Cn9wc50A4m` +
    `x3DCyYJkvVlNY72KWstV0LlctfIutag2RvinU6/Tep4N5+dW+JohELScXnSKjCv9vB` +
    `MWNbTDdgSB2oJDWAH2Lj8MaP2e4vrrKnpV2GdlzuaXpKJqKkZQQKA24BlEJxNlzyTK` +
    `iTqwVUMvdrtR8bIF/Jl0vabnzzhZ7nQLUWjHSpZre4AD0zp9nlGU6vB70gavsmH2lG` +
    `PMnOZ5/3xtGxEYI7BrrAh6rFBDcFiPKUANCvtNx6my/WKLRkzK82dDpR6u1IuNSard` +
    `q56K5Wfl/f7oxL3hOZr6FKQ/A99JC1CUeH7opWoXdzGOo1Dt7LoL2y5yPT/0VT4RRk` +
    `HsRmp9imI3TJJmloBRGnpp4NlO9Q0rFAbIV1+xnlt+xdP2EA60mMb4RB3F1CgbI1/J` +
    `CVOVeSgxXmB74CE/NuIDFGjpHqzAQ2mSaPV+Oen6LhMDFOiuwMYRCiBBruGrm1rphn` +
    `oxio0gzUB/67kjQ1tzfOQZczBGsbZbSQixMgdwWHHTUpv+DJCvrdQClZ494Opg3mHS` +
    `SdhFKEkNjwBFobZXpYna72FiYxTozggw8mwPJYb5IQgwVgfi6jG4Fa9IQSOs4HAqXG` +
    `ojRIbODpCHjTEORvoLVnrouDXh5xqgKfSklYituraL0srZKKwxGhqrGiwcF8WGh7P1` +
    `ZdijZ4LUBoMCE18Hm0AFatJU7nNr4LmHoDqNWcW4ip5rwI/1ErKVAkXL83gr2TuieX` +
    `sRHNyI1bRZNChYlgNbcvgpW2VizOlC7koVVVvw8dDaPUc/CSXENJq85hktYcpp3lns` +
    `UE93JYUt/XvOGYc5ESKbkl2VpBqq1raaympyMuNGjOeEkxxMJaUqiTQYNEa8ZaXkrG` +
    `guhlsRY8mLj6yc1hJqvvWT1o2KUn685JyUsthANpZ0RUCwctrQvyVAV5dq7maAYvtJ` +
    `9QMngv5CrDvdatnQYpJxqS9iD7i80Q9qNqZ5gk8xWs672Kj+HRfV6mZyMPKesykndZ` +
    `nwQOKifn7I8KiapMgWQuV/lTgVivemD8Q8Kwrrzr318HEG+ZJnkrKyxeFd1bll4d1i` +
    `b49Fs7BWh1usqRzPYETkmpASdAVXWK14N32jRH2mZA0j9vhAxIKVgq6IdcyVW/IfVI` +
    `D3nKhnqYovFSBnBKqyrN7n9/AGMBgtpWSltt9c22/zaF3bMxl1e0hVJaw0UZmEdVbJ` +
    `cJ8LwGApCDwWtPz51tRs/2dXPrgb3CwFOZBslqN2lzGhHyR7vvtoSijH/SG6CJ7d+q` +
    `3UZ7f/O87WB4o17R9pAv7SDqiqH9d0QOcMXRTZxgZdtQUix+jk3Lwvsg3hF8xOM0D9` +
    `f8LDpkatF/uh1UH77L7WCjzB00ct4sTcoFc2PZjmH92qqrS/DZRpdb2T8VIGag2ed4` +
    `kwAP4o2eLUEqlvKSq/dNA+/2qpZb6AI+45WVG2PEhZmijZktRA6R7z/PDgZPUiTvmR` +
    `PB6kHw2HmLdLKmd00D67I7QCL+GIBzJnq1N7KTcEr+mMUVa+iC/e6mul466orp1e0R` +
    `NGgyv7QqUU///Xf3x4fwtZnoOlq9uV2ZKBAYBOyKuT4OEB8CBN2fPvB5aTUylIQ4UJ` +
    `LXPI1PFDEg7rbKN0kGw6LZQOVHQdPPckPVRX/1WgPpnB5rYd5uxg12vBXVO9MweIV9` +
    `0qqwLG7wp2l23vS23aq1izq3VcA8N9MKgKn3sw0L29MNBUXwQMjBXXhIHm+Eow2Cvy` +
    `XmDWfsi7g/6xKrW2g24KnF1B38sKNd027B1WXck7x7NCo363WzodcyziXYYdifmL2d` +
    `UX94sMPBcRb7fvRrUhUb0+04+JivAVQbF9u+vqqKgYvx4sDt5huh4ujnaYpk45SJkD` +
    `m0Djzrz+PQE0L1v0peTn/VvI+vp1Rcn6viKz9i/iSSkPr1u2VyjFkVvQ7bP3peSUCJ` +
    `gymBFO9nOfM++KtPqklHwDkswXRSYJrGlRwIio52OSAy0la2vVbWpTsU1tq8oHH9h6` +
    `v0hxZH9W1JW7T+SEe5OZFMXJWbooOm+hXgDDJ+5sfh9uL/ZA51nuRXxw/IR21al7iK` +
    `N7TtlVQcQpw/VWYQ4+6vixgH8vyXKXCS44Za/h6OOXRi+MNeUo78BR5sWZL8RX1Vs8` +
    `X4y7/AN3mbdyvhB3Va8Ivco0fnvowXrL2T7/1Kost3p76sodzN60K8zt7r76che/Vq` +
    `W51dtTZ+5g9m6/nLv/oGbp3fodZeFziry6VteT01R1xo6M5sgGv6s9fsceG68YPDaT` +
    `g658p1Kn8TPSw3rv7kejhvqTuaNt/HjU/OK4s/jZOgstR3MqO4uof6Qa6XHnPxCxLH` +
    `bgVvka3Xuxo+13rgc8Ibusxu/llp2AMqRfXN4o+Js8t+6+/XMl+wOuZM3F6h3lZCxZ` +
    `A13nrVf5dtwPTTeemDP1AKiILpw3nYLb86J+dP7kqIdcfYbkRef7FvASOcnJNypeNI` +
    `XLi+utE7tYnVgsaqIfszk5RIbqrSePalw2ebovQXZ/TeFmxPKNbgxuzB8R+U8AAAD/` +
    `/xqE2olVRAAA`)
  if err != nil {
  	panic(err)
  }
  container["templates/javascript.js"], err = decode(`` +
    `H4sICIg4slwA/2FtRjJZWE5qY21sd2RDNXFjdz09ALQ77XIbN5L/9RSd2VtnGFGkJC` +
    `e3OcpMKrGVOl8ljstK7V6Vy7ULzTRJRCDAABhR2kR/7wHuEe9JrhofM8BwSMnZXf+w` +
    `OEB/Av2FHswt0yDUEuawaGRluZLl2ixH8OsRQKWkUQInQi3d4MXRw8URIdR43Qyj8A` +
    `WUCyYMjiDDmk7BrJS2sFAaalU1a5R2skR7KZB+fnv/ui55PXLkMSXNa09Zo220PITr` +
    `GBGnlbUbM5tOjWXVjbpFvRBqO6nUevpLg4aomumfPv/888+/OJuu1PbEqpMl2hO7wh` +
    `NW2YaJE42yRo31yUJJe7JdoTzh1pxIZU9qXHCJ9QmXJ5UxTt7KmFRi9GKNYaPVBrW9` +
    `zxTYclmrLYn/Uq03jcX6yt4L7LBkI8SI5t8G9D8z0WDZEvsoNc/Ovjz//GxaMVE1gl` +
    `k8sXhnT7a8tquTLberk5/ZLTOV5hs7PT87Pfvi+X88/0P84ZRbov0J7+xfCCXVkgiN` +
    `gdbniv8dvYrTKRnAkt+iHENjECpWrbCGislbZtzeX6O1qGGDeqH0mskKPR4Kg2OoND` +
    `KLIHEbcI4A3Ap7AvMjAMgkmoSZ335zU+XQ3LwzGk8/2E1ZeIBiNLoIfEid79iai/sh` +
    `XsnsIL8Um4yixLK4VvV9MRpD4Sxp4WYThpWStJAE72TxduHGyuK8LhxgAHL0aQvCms` +
    `MxFJs7KOA4kTsSXqPVvCLlI/YamWk0krhu8xzpYJQBeuIM46McdlIJZsz33NgJq+vS` +
    `PXkvZnVdiZ5jjKESxtsKih1U83GxIsHX65SzXj+Zsca1usWON6GveI05tkf0CjliBc` +
    `EUHYpZqe0QipNkEKPmhl2LQT4oJmG2hjkU8XfR4qJ8ImqR6fS6HoqrNFMiPebqDEPT` +
    `TB86cBxGCJN9HK/CMIqfSzGOpp8dwWdgttxWKwoh2y1KqPligRqlhVuOW3MEn02PKG` +
    `WFZ5jHmESPINkaZ3DdWKskvH5FbsJx+1awe8GNnUFBf4pxGL9CpqvVDArjfrTjr7jG` +
    `yip9P4PiWqutQV2kq/ZnYpXoRDhZ+G9n/Ci4NFmsmKwFzsiVHcaFmyL7L4nyDXDplY` +
    `pYTq+wttNgD05Nr6CBZ06cqbNlmghYfn1QwBywvAl8IK5sedPRB6iY8bgz8C65hTWz` +
    `1YrLZUqyVYIgnAodXUfbmYzIxlqTcQtk3t98yKavNbKbbqDGBWuEnSUQlP54XSMou0` +
    `Jt+qK4uQFZvLHnskRTfIIoD0fd/7QkG6w4E8Dcjhq4ZgZrUBLsCsGgwMp2CxWXuLOI` +
    `do2LzLKKTs9KrddM1mUwtWIMRTG6cKxR1sAiwNGgtB1tb80JYSrTnLnM5xlEuv09Yw` +
    `meQOmjyNdPTLyuZX94oarGpKMPQ4LS4IN385A2Lm/RJ7o0ftv7DY5hIbs6E8UkAlMs` +
    `R4k6yj8wVUYCY/AFKnF/cHVHgmQtq1YOryyUJANq+Tr4WHu9Q0MZ+FapWimfdPxIT+` +
    `o1l99djWHN7r678sK5MoOCkxsL6d1YsK4KQDHhUqL+z59++L5N5ewuVl8ud3GUvuAg` +
    `gO2KC4RyYeCruecGz55lZUtpXaFmRlQnGPiqpRcXa2HgZA7nXr9Qwht44aklQIG+37` +
    `VWqM8IlUxyi8CEUFtQEmHDlug32UGTGlY3smLWZSa3A55QUMDe2YlAubQr+Aq+/Ggd` +
    `gt9bpUAouZy5qsiGIof++eUlLkbwCsvTMSQsT+Csg0zktLrBi8TpaW3a+Y53umsHuT` +
    `wnBYpnKxSCby6KxAF88jZ0FJi0Bd6clHVVXpfJvZ3p1NA6w3K+OlAymW/vX1LN84at` +
    `sSw8iaItQjn9cMmGwxxOL4DDC5LGi30B/Pg46upRSxTv+YcxnJ2P4fzLxDP8PvPlEv` +
    `W7HX/wNML5p+ZmQ8nE+xoV/MHronA+zgVi0buQYEwbYpVcgrHNYtEy35q/srrWY9ga` +
    `76h5oIyYW7w2qrpB689vYTKRtVrXY6iZZV24+YSyb8zi0X7J6n6FFhoe3GJszYT4lv` +
    `919eObibGayyVf3BOg5zTrMGYRLy3HvdmFULPVbLNBDVb5I1UQlhlgEpSsBK9uCA7v` +
    `2HojcEaxeiPYfTGahFmYw7WVL2MycZNjOPPFcjdzQP99xUsYj3mqw7rogvp0CgwWuI` +
    `VPl0JdM2E+JVWsZtWN24uNVhUa0qaGqtGuoDOWjoFq4QEEu0ftD6OvGs1IApjD2eTU` +
    `e8TyUrCNcf562o5dOQq+BCQCMyg2rKEcGiJuQ+Wg28exU3CDzCYDhsulyCCYrNU6hX` +
    `Bx5pR+Sryz4TG66fJSa6UpaL7/EEZirfkdF5hOaGQ1lSBXWClZZweXWzr6d+695qT4` +
    `hmmDr6X1szCFfz9tXdlgtQvxxwhBZkwQL+DsNG6hxyhOKWAarGIWiAdDLikAzdrZsK` +
    `N2xU0rpgHdSHe8x1vUbiOXaA1UTNBBhAbvCVvJ2lnxpqa9sSt0Zhs2j7DquLtcuu6A` +
    `+x2sgIqsW+/zbtU8lbdaLTXm/ZdwiiiLQLoYZdG5t9xlNB+3QlgWUYjH0KIptivrbe` +
    `69968PrrRyP10yizb6orPhuAPTKXBZaReoQ+noDYEvnOlSvU2rI5WFBZfcrNBXfC3R` +
    `4zmcpel7QBJj1aYt7yhChIUrRiHpbEOJUZz+sdgtjQ4gtAsI01S3z+Ds9JRsJ5Ajj0` +
    `H7E1+jamyZ796YYE9HXRFIZwbnQHv7kl1dumZcOtii3Ym1WcIndP6NCgdvnGwaswpN` +
    `SS/SdArfvf7vHy5nIJS6MdAsxX3rSptwhg7PVrfPkd6abcohZwUHfTx3RI7BzVwEd3` +
    `NUX1xrmH7lV8YbXl5CGKsvjtKzU758O4HY7XoQqi2jTrsKJc6ZFV8khXq7zmURi/qH` +
    `ffTm85QgnaQ8ZoL3MIYv8o1swXYd9OAu9pbDr3p3eovkvRmRtTdZCOgyl2tVl0UK6D` +
    `l0O09ZNUZwKJVuH3g9oni14suV4MuVd06a9FEoqhA3nFmXF6ANY+NwrHDe4R9C5hmH` +
    `7ElJJ/x26cX/9rkn/FZy6X9FqZxvzl2adaaYZrqY6xJh2lw3wHqWiRBTXso+Zr1MPJ` +
    `fpdgWbZSKGNfbr/7d0/ct/+9XJ9zD6m9uJeDJ2g9HC/OHVp+xZOkIhLR5krVouBVJe` +
    `fUuAgUI07rRYiFvSznU1Q9ifOJOcTD1HFzkjR3ooBzm40mOHeDLaEg5xp0uDsFY1Qm` +
    `W1OAJ4X4SNotO+3x765TfFjbmtKD4cCD7BtV5x45zLVwHHULzybZfo6gHsUuZQlzIF` +
    `SvKJg/jQxQAfnV5x0y6I99BLGWNCmkPibArvKXTwIRzm7v2Ny/xX5J2DLh68EBZc4B` +
    `iYttzYMVhu3aO4btZ/jYPuaQwbZbjbttSRsCxoC4hPMZo4dOLHBYZZT6RXFvjBAOGQ` +
    `egBuLFIg9vm809uNwzGUnhrlrlRu+BqKF9tr/dV7KsOymWMoPhQwg6J9U5DEVN/c7t` +
    `Jiq/Z8DidncR/bDq0YaInIdMqF5eEzRWDVLe3Eqit38ClHvfOhxC2t8RtV997zjQGl` +
    `1fdjkPqpG9vf37aApI111GL2lp4dhhcHSiIJQId5HzBofve9QNu1DwB9q3CDnByd13` +
    `AMUreDvzSo769cX4yy6x8oLL5EIb45bEQHcX86ZF7TKSgp7n2f1htUsBG+AG4/NV2/` +
    `/FEZ/9VmepD7q6fV33VafofTCtHtokc4WL5Ds1HSZMamrn/OTcx3+fxRHOagrn/uTv` +
    `fpXKjEQsaimTxhoSth0oRFR5kufbT1VntKHsw5t6gNLcEsbW798PYV6WRVpQQECN/o` +
    `OkzMR9GWVpaKCXNi3G+vWqz2w1m1twlFceGTukA6n4saImBMfdkhty0doc2ClPgoNV` +
    `oFC75sNIJqnIluqbTyE86EN5qr9m0GFWWqfQHgXzP7HsF87nIynbH8CZEYAzdULdPg` +
    `Lw02vnf0Yn4O//c//wtS5YwCj4+hTScxok/kHqHxSUdjV5RaeVKDJEJ1yg3YFbNQKa` +
    `2xsl9nL3IcydAqdFvpBmLjLgMNKkGy4+4P5mCJogGWtel3EsOeSxsdu6wQCTGcZ2+S` +
    `skBOO9slijFw4zM8/XpD1XUWvP0/Sl8RLn9J4V3fm6WSn4bV7HUMgPu+X2Kp6duInt` +
    `W6UyLJmry6SHJIksA6R7kkeYs2hfEE1Wt/bSVxIfyh6LcRb13P7ikqx4wdSGavXcJL` +
    `auIxhqLbu5RyryxzQcOTOtwx3Jfb+8vpJaetzOXOZaPDwo5k7a89i/QulDTDgia82o` +
    `okGRuQv51NhOh8z3W5XFZdsxukmEJ5hEuoUFvGpSvcDZS+JIdj+Ho0sOuaq7MD207T` +
    `xaC5aK7ODyOe70V8fhjxed/SPok9o3i8yLfOle5Rmd1XqlHavTPPD5ufa9Qj4N1GMO` +
    `krOHatbjEBISGTRwjnRdfncUc0+O23NHqF/s+zZzlS9CiC7qLnizmcjxLAXPnc43ZW` +
    `YGf6/PD082x6YDW6AoNCZV+WkNZPZ73RR/j4f7036wm9s/M/PUqxr9hhiudffPEoxf` +
    `5K7qH4kD1F3H3RSnNFB2XiT9XR4bCVEDx/hCAt0UcRfP4IwdMnkXsYCI95icY2G5S1` +
    `C6wt5kNXz/WP0WUvryevPpNXd+X+itJfBXiHphG2rSvbKwJh/LHK0QODdtBd2eMbRr` +
    `/C0uXiWf+1yUMrFQUEp8jAC9mss5kxAsHX3PryL5bOkzW780J3nYmHPpdUt+GWaMZU` +
    `KusuyixUI+tih+wuyY8qooZKJD8/VKN4Pk+pUPYlDaO/qesiz5HxBlu4UrVQ2olmgA` +
    `k6q93Hmqt3Qggdpbzi4rISTY3GV12jfm4I6disVCNqd9GgroFFzlaBT/bALSy0Wmds` +
    `vx4unIZvpViZeOxOk/0QkZ13oQWr62IMvTLy4SJ14Z6/7HPjPc65zzvreJPoe4oOiX` +
    `vmEzv+2bf4FTNvmUaZVXFtHyWn5o1rsLPS4flOSVGLd2rrKRc9gNhfcfw3DuQi8658` +
    `cfaUia1kb9gaB/oYiQt6Hr5YgK+zsRkUk0nRLjPA+z7lMRR/qMW3/m5Wvx+bG82AnL` +
    `f73s+3l70SaVITSn5HWxxeBuGcdpSZXN8GBm0uDVIZwtOjVGiaROTx3rAVD1Z5OUro` +
    `7r1lS6HYXdDfY4e7lkhZgGeOvE5OKV0D8PeaXLcIKZV/xJh+rzlFfnl1kT09mgmCUV` +
    `3kZ8fHo+b+uLkvcu4R9+A54ndlv9+T//6xDPjROfCfmgX3b8XDv35Xe/kw39CPDVR7` +
    `K9aH3dq1ewGhNij/gtdX7gLYwAU14y0IWpgy3CqLd7uUJBLZVc/bNkmGF84/vr18Ew` +
    `2m9678oSVTCZX3pXfpvPz+x6vLSMiL1gjhH7Nyk4xYYuU75zmbNRrDlvsY9drk/tKa` +
    `u0FEUBPXW+4RxP4b/YRcJ1SLfZFdBgtmTJYJU3BveGHqm6S1QiM/tbBV+gaYAbzbYE` +
    `XFekkzrlca7rrHxjATwtm55Wv099kcoZ1NDavpb8FcuNes4ToYFHQw8AHYTYfWFP3d` +
    `aLzlqjEhHvuPKCK674nFPpOH8B9OHIKIt827TyE2vis4LPCmbQq+T+QcpJxLF5V7RK` +
    `Vc4CciDejgdnGvEo5un8njSjx5YzItnoy1q0bvPkGqUHI5obvV0L9/8BEXG9JrDPRw` +
    `4LSdXoIAv9QD0D6+hdu1O5ffC6GYC787Hhtim6+ltqYYTfxtJVKTspjPJe6DCglcbh` +
    `oLC46i9h9E+Uu7nqt3Jf53WvtwS9mJuQumNEdpXZ+vWjG5zDGOIA/R4ayTXxkLg9Op` +
    `y4IxUVkVbhv6+HB08Bsa7L74SDJd/IonfB3z4G3KN0XHcPByxqErGQP3K4brNQc2hg` +
    `KzWxhDlzgO49fpXY/QDKJqqLlecxs/9xjK8Eks76fw+DFU/3MQbzIxzh/Bznzc5382` +
    `nzTItsHx0ShG2yiQyf6GOVkci6OuBRxW/pHVzpY5OKGS7gOYXOfkOO0czH2qCwwWGs` +
    `0K/Hs5Sn1GrRHCF2YGSpzAcgIvV1qtcQSmMVQMgWp0pBSvZya3BLv7uT6YxNX1jvTO` +
    `sy464aNDtS/Q0g/kwmXE13VZpMPui8aRT+yoFkf/HwAA//+No1zlTz4AAA==`)
  if err != nil {
  	panic(err)
  }
  container["templates/style.css"], err = decode(`` +
    `H4sICFU4slwA/2MzUjViR1V1WTNOegCsGftv47b5d/8VnIoOiRHbSXpZLwoGbMtuwI` +
    `Bdd7gMw4ChP9DSJ5sLRWok5UeL/O8FX7JIi7KTHIrmQn7v9ydmMUV/+/t/Pn/KEeX8` +
    `uW1QK6FqKfpTDSXB6P8tCAISTReTyZKXe/TrBKEDzRKUAoEqzhSSQKFQhDODbLAMYF` +
    `bhmtB9jrKvfMkVR58549kVqjnjssEFPEwXkyNUslor4K3MrlDRCkk28DDxWJL8Ajm6` +
    `/djsHiYI1VisCMvRB3tscFkStsrRtT0XnHKRow0WF7NZxQWsBG9ZealhS1w826NHON` +
    `xcPkxeJhPnhouasNmWlGqdox//8LHZXRpHdB7R1vaZfQeV/s9aNqj1i+ZOWNOq/6p9` +
    `A3/MFOxU9vNVcMfaegkivm2wlFsuyuxnI9updXN9/b2xiYsSRI5umh2SnJIyYbpBmw` +
    `lcklZ2zlvyndbT+M9hLLmBzLawfCZqNoJR819GwD0f3EBtnLuYooLXNWdIqj21WTZf` +
    `kpUxq49/q017mUzmS07LA3QLOku0IFpauI5WWx/RX3t6WWNKY/D1/Mc7p9GcFJxZOK` +
    `E0R4wzeOj5WKMhtHZy3ZFvQFSUb3O0JmUJTN+VRDYU73O0pLx41jdSCf4MiUS0QJ9h` +
    `H7/v3VHC4H9cJ7hB937QroXSlyMuS4QtgnM7+j2SoJBaQy/LkclHnyG3b8kQX203Qb` +
    `XNllwpXjuswcoq7i697ruvIBvOdFF7A+SWqGJN2Eq3lC0A03Xb1gxd1HxJKFwizEok` +
    `+BZdlCCfFW8urTGdoysKNs8o7GYlEbYX5Y6PhmBKVmxGFNQyRwUwBcJqtCalVaThkl` +
    `gqvJSctgrQ70jdcKEwU5qF4k2OZvf39/fNLgJRqNQw7GUyWbZKucTq+tNt6NDb4wps` +
    `mQT18N6Ijeau7joz45iDS6LiuLeYIxFtBOGuH57Ve6zJM2E1MWp2PsrXupwMN97ggq` +
    `i91uHuvG7tWFS8aOX7WJRE4iV1BRZzCUZK8SExSorbc0bIQDkc57Dg225maIJHWxu+` +
    `+OkW76UvmNdVhS9Ivo246VI7j5VTLszrHt95hW0wNiAUKTD1+aZ4YxELcE05kmVHRI` +
    `0JQytBTPeam1OAq0FaAf3vTEHdUKxgZg20VYBuKuGrIcTCArDMjc+zuZH0SQgu0Dzr` +
    `3z1ypgSnA7fAFJpn54S50/tcRftrFsV73irktomkCcNGpMwYNMQn2YHLr95pWlB+YB` +
    `8n/XdVdVPdVA8duRc3wMCBQlytQgIXXBtdTBGwEvEqTAqfJV5ikCx9Nc7NGdwqbn/o` +
    `gHSn8fwplKD/JrBF+pcnzlYjv2TfhuQLxXsQaG4On3kJ3S+ZqyyPHTnW3x+wjKwBLH` +
    `1/wPpEcSNdV4wQHag3cT7YwREMkzt7Z0exBFrlOqKpSeTF/rUV2HxZDMj1sGPBZ8y2` +
    `lDpesHPwgFgLOWAa7w/g6fuB3I1T1ezCPlhd/gbhe03ynpep/yBSRWn1BFgU6+jyL4` +
    `JvJWT99eEHvzD3eaWyTMNCbCsmhW+hIYXVIUVhoZGfj/zpneyiGri5F+nXd4nwR8r5` +
    `1qFWkP7Z0+VJ8aZ3/CJgQ3gre1c/wU71jl+h5hvoXTxSwMyEaDE9GfykDtkRzrAmmZ` +
    `1Q00VUKYZlslr0z7i2jOVJiqfDjnCsUlqOw4gpjQ+TVBoaUzg3J2ksPKYysUgTGfBA` +
    `rkZp6bPVdJYgV7te845MHU9SLeGRM9nWh5HyFRrAqjs+EbaiPShmJa+zsCF6Fom+6M` +
    `AhjROTILHQkMJpkqCw0EiGUTYpw0AtheKrFYV/2miagVJCwe3E0R9oJQj9gT4QzSBs` +
    `PpZmyq4E6ULZjd1ufXbrn2+0H/yMelWs/doyPgdwocgGtAr/IooCOnHOYqo/C6W7/s` +
    `mLY0K6bGt04tzrRX7tOGB8EXwlQEoU7AAu/WI7ojhH4PSq0F8L+ivJkcFJARb+LgnG` +
    `NWkBGvwe/p0jkyI8xiseF4fFDkktyca+HIx/mB/q4Xrutjv/2OGeM8YUetHlh5HUtU` +
    `YYwojqoOkS1Ff+2/dNHwfpNqpZ+y1PGpMJN788AqVZ7xkkeDo7/SHZV/l8Hd+u5bCe` +
    `/jvRa9N74PI54t7f+5G1z6kvB7J5J3Po3eWsbLONFyP7IsHFPgpxd/+mOI+GuGPtPX` +
    `i4eF+Mj3Q+S8l3qDke5ECdV0Y6oP0m4Z4f2RO1rhgeEfnHpkEKoDTukFeTOYOd6ib1` +
    `Yjr02Hdz2X/z9Zuxr6RYoL9/zWDvz3T/7G7eo29PvWx5aXYa+lM3NEMkN9LCY4xjhl` +
    `JwOucRbDFF7vHePmWF7hk13P/vHTDazzpTUcqoQP/e9tBvawc2ieB120NqykbSE2wO` +
    `O8JJPkO7QAA8h0vqMSWGx1uFmbSxAKmwUP2MdPPYPwIvpqhlGxCkIhXYVXiuWzM5fm` +
    `1y11ZZd/j2Pdsx9q3DHT8xJfZZIDrRXELoK+t3vExD1l8EqbHYZ0OwJyg4K7FX+cQ0` +
    `GTDoVK0dKqHvoIQvDOytnSztDs1KS7A16E9R1/LXrsTCY4xjyic4OYyu//+Ea0DBKY` +
    `veDD+6PyW/stMduTHpnB8qgW4rYf493fH8S/6ws2KvHDsg4ddxn5x/6rXUUMMokwLg` +
    `Q4A/3DxDaEQx1CYDYIifaogx3FCVZNNrTjlT61mxJrS8gA2wyysDd9ttDEx86Pg/o4` +
    `5zjjemcUlD69XxXy3HhPKyTFqjYSljzuF7ri1DuAOm3HVfAcArXW+/BQAA//9f15iO` +
    `1CQAAA==`)
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