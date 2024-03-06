package distribution

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/osbuild/image-builder/internal/common"
)

const (
	centosGpg    = "-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: GnuPG v2.0.22 (GNU/Linux)\n\nmQINBFzMWxkBEADHrskpBgN9OphmhRkc7P/YrsAGSvvl7kfu+e9KAaU6f5MeAVyn\nrIoM43syyGkgFyWgjZM8/rur7EMPY2yt+2q/1ZfLVCRn9856JqTIq0XRpDUe4nKQ\n8BlA7wDVZoSDxUZkSuTIyExbDf0cpw89Tcf62Mxmi8jh74vRlPy1PgjWL5494b3X\n5fxDidH4bqPZyxTBqPrUFuo+EfUVEqiGF94Ppq6ZUvrBGOVo1V1+Ifm9CGEK597c\naevcGc1RFlgxIgN84UpuDjPR9/zSndwJ7XsXYvZ6HXcKGagRKsfYDWGPkA5cOL/e\nf+yObOnC43yPUvpggQ4KaNJ6+SMTZOKikM8yciyBwLqwrjo8FlJgkv8Vfag/2UR7\nJINbyqHHoLUhQ2m6HXSwK4YjtwidF9EUkaBZWrrskYR3IRZLXlWqeOi/+ezYOW0m\nvufrkcvsh+TKlVVnuwmEPjJ8mwUSpsLdfPJo1DHsd8FS03SCKPaXFdD7ePfEjiYk\nnHpQaKE01aWVSLUiygn7F7rYemGqV9Vt7tBw5pz0vqSC72a5E3zFzIIuHx6aANry\nGat3aqU3qtBXOrA/dPkX9cWE+UR5wo/A2UdKJZLlGhM2WRJ3ltmGT48V9CeS6N9Y\nm4CKdzvg7EWjlTlFrd/8WJ2KoqOE9leDPeXRPncubJfJ6LLIHyG09h9kKQARAQAB\ntDpDZW50T1MgKENlbnRPUyBPZmZpY2lhbCBTaWduaW5nIEtleSkgPHNlY3VyaXR5\nQGNlbnRvcy5vcmc+iQI3BBMBAgAhBQJczFsZAhsDBgsJCAcDAgYVCAIJCgsDFgIB\nAh4BAheAAAoJEAW1VbOEg8ZdjOsP/2ygSxH9jqffOU9SKyJDlraL2gIutqZ3B8pl\nGy/Qnb9QD1EJVb4ZxOEhcY2W9VJfIpnf3yBuAto7zvKe/G1nxH4Bt6WTJQCkUjcs\nN3qPWsx1VslsAEz7bXGiHym6Ay4xF28bQ9XYIokIQXd0T2rD3/lNGxNtORZ2bKjD\nvOzYzvh2idUIY1DgGWJ11gtHFIA9CvHcW+SMPEhkcKZJAO51ayFBqTSSpiorVwTq\na0cB+cgmCQOI4/MY+kIvzoexfG7xhkUqe0wxmph9RQQxlTbNQDCdaxSgwbF2T+gw\nbyaDvkS4xtR6Soj7BKjKAmcnf5fn4C5Or0KLUqMzBtDMbfQQihn62iZJN6ZZ/4dg\nq4HTqyVpyuzMXsFpJ9L/FqH2DJ4exGGpBv00ba/Zauy7GsqOc5PnNBsYaHCply0X\n407DRx51t9YwYI/ttValuehq9+gRJpOTTKp6AjZn/a5Yt3h6jDgpNfM/EyLFIY9z\nV6CXqQQ/8JRvaik/JsGCf+eeLZOw4koIjZGEAg04iuyNTjhx0e/QHEVcYAqNLhXG\nrCTTbCn3NSUO9qxEXC+K/1m1kaXoCGA0UWlVGZ1JSifbbMx0yxq/brpEZPUYm+32\no8XfbocBWljFUJ+6aljTvZ3LQLKTSPW7TFO+GXycAOmCGhlXh2tlc6iTc41PACqy\nyy+mHmSv\n=kkH7\n-----END PGP PUBLIC KEY BLOCK-----\n"
	rhelGpg      = "-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmQINBErgSTsBEACh2A4b0O9t+vzC9VrVtL1AKvUWi9OPCjkvR7Xd8DtJxeeMZ5eF\n0HtzIG58qDRybwUe89FZprB1ffuUKzdE+HcL3FbNWSSOXVjZIersdXyH3NvnLLLF\n0DNRB2ix3bXG9Rh/RXpFsNxDp2CEMdUvbYCzE79K1EnUTVh1L0Of023FtPSZXX0c\nu7Pb5DI5lX5YeoXO6RoodrIGYJsVBQWnrWw4xNTconUfNPk0EGZtEnzvH2zyPoJh\nXGF+Ncu9XwbalnYde10OCvSWAZ5zTCpoLMTvQjWpbCdWXJzCm6G+/hx9upke546H\n5IjtYm4dTIVTnc3wvDiODgBKRzOl9rEOCIgOuGtDxRxcQkjrC+xvg5Vkqn7vBUyW\n9pHedOU+PoF3DGOM+dqv+eNKBvh9YF9ugFAQBkcG7viZgvGEMGGUpzNgN7XnS1gj\n/DPo9mZESOYnKceve2tIC87p2hqjrxOHuI7fkZYeNIcAoa83rBltFXaBDYhWAKS1\nPcXS1/7JzP0ky7d0L6Xbu/If5kqWQpKwUInXtySRkuraVfuK3Bpa+X1XecWi24JY\nHVtlNX025xx1ewVzGNCTlWn1skQN2OOoQTV4C8/qFpTW6DTWYurd4+fE0OJFJZQF\nbuhfXYwmRlVOgN5i77NTIJZJQfYFj38c/Iv5vZBPokO6mffrOTv3MHWVgQARAQAB\ntDNSZWQgSGF0LCBJbmMuIChyZWxlYXNlIGtleSAyKSA8c2VjdXJpdHlAcmVkaGF0\nLmNvbT6JAjYEEwECACAFAkrgSTsCGwMGCwkIBwMCBBUCCAMEFgIDAQIeAQIXgAAK\nCRAZni+R/UMdUWzpD/9s5SFR/ZF3yjY5VLUFLMXIKUztNN3oc45fyLdTI3+UClKC\n2tEruzYjqNHhqAEXa2sN1fMrsuKec61Ll2NfvJjkLKDvgVIh7kM7aslNYVOP6BTf\nC/JJ7/ufz3UZmyViH/WDl+AYdgk3JqCIO5w5ryrC9IyBzYv2m0HqYbWfphY3uHw5\nun3ndLJcu8+BGP5F+ONQEGl+DRH58Il9Jp3HwbRa7dvkPgEhfFR+1hI+Btta2C7E\n0/2NKzCxZw7Lx3PBRcU92YKyaEihfy/aQKZCAuyfKiMvsmzs+4poIX7I9NQCJpyE\nIGfINoZ7VxqHwRn/d5mw2MZTJjbzSf+Um9YJyA0iEEyD6qjriWQRbuxpQXmlAJbh\n8okZ4gbVFv1F8MzK+4R8VvWJ0XxgtikSo72fHjwha7MAjqFnOq6eo6fEC/75g3NL\nGht5VdpGuHk0vbdENHMC8wS99e5qXGNDued3hlTavDMlEAHl34q2H9nakTGRF5Ki\nJUfNh3DVRGhg8cMIti21njiRh7gyFI2OccATY7bBSr79JhuNwelHuxLrCFpY7V25\nOFktl15jZJaMxuQBqYdBgSay2G0U6D1+7VsWufpzd/Abx1/c3oi9ZaJvW22kAggq\ndzdA27UUYjWvx42w9menJwh/0jeQcTecIUd0d0rFcw/c1pvgMMl/Q73yzKgKYw==\n=zbHE\n-----END PGP PUBLIC KEY BLOCK-----\n-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmQINBFsy23UBEACUKSphFEIEvNpy68VeW4Dt6qv+mU6am9a2AAl10JANLj1oqWX+\noYk3en1S6cVe2qehSL5DGVa3HMUZkP3dtbD4SgzXzxPodebPcr4+0QNWigkUisri\nXGL5SCEcOP30zDhZvg+4mpO2jMi7Kc1DLPzBBkgppcX91wa0L1pQzBcvYMPyV/Dh\nKbQHR75WdkP6OA2JXdfC94nxYq+2e0iPqC1hCP3Elh+YnSkOkrawDPmoB1g4+ft/\nxsiVGVy/W0ekXmgvYEHt6si6Y8NwXgnTMqxeSXQ9YUgVIbTpsxHQKGy76T5lMlWX\n4LCOmEVomBJg1SqF6yi9Vu8TeNThaDqT4/DddYInd0OO69s0kGIXalVgGYiW2HOD\nx2q5R1VGCoJxXomz+EbOXY+HpKPOHAjU0DB9MxbU3S248LQ69nIB5uxysy0PSco1\nsdZ8sxRNQ9Dw6on0Nowx5m6Thefzs5iK3dnPGBqHTT43DHbnWc2scjQFG+eZhe98\nEll/kb6vpBoY4bG9/wCG9qu7jj9Z+BceCNKeHllbezVLCU/Hswivr7h2dnaEFvPD\nO4GqiWiwOF06XaBMVgxA8p2HRw0KtXqOpZk+o+sUvdPjsBw42BB96A1yFX4jgFNA\nPyZYnEUdP6OOv9HSjnl7k/iEkvHq/jGYMMojixlvXpGXhnt5jNyc4GSUJQARAQAB\ntDNSZWQgSGF0LCBJbmMuIChhdXhpbGlhcnkga2V5KSA8c2VjdXJpdHlAcmVkaGF0\nLmNvbT6JAjkEEwECACMFAlsy23UCGwMHCwkIBwMCAQYVCAIJCgsEFgIDAQIeAQIX\ngAAKCRD3b2bD1AgnknqOD/9fB2ASuG2aJIiap4kK58R+RmOVM4qgclAnaG57+vjI\nnKvyfV3NH/keplGNRxwqHekfPCqvkpABwhdGEXIE8ILqnPewIMr6PZNZWNJynZ9i\neSMzVuCG7jDoGyQ5/6B0f6xeBtTeBDiRl7+Alehet1twuGL1BJUYG0QuLgcEzkaE\n/gkuumeVcazLzz7L12D22nMk66GxmgXfqS5zcbqOAuZwaA6VgSEgFdV2X2JU79zS\nBQJXv7NKc+nDXFG7M7EHjY3Rma3HXkDbkT8bzh9tJV7Z7TlpT829pStWQyoxKCVq\nsEX8WsSapTKA3P9YkYCwLShgZu4HKRFvHMaIasSIZWzLu+RZH/4yyHOhj0QB7XMY\neHQ6fGSbtJ+K6SrpHOOsKQNAJ0hVbSrnA1cr5+2SDfel1RfYt0W9FA6DoH/S5gAR\ndzT1u44QVwwp3U+eFpHphFy//uzxNMtCjjdkpzhYYhOCLNkDrlRPb+bcoL/6ePSr\n016PA7eEnuC305YU1Ml2WcCn7wQV8x90o33klJmEkWtXh3X39vYtI4nCPIvZn1eP\nVy+F+wWt4vN2b8oOdlzc2paOembbCo2B+Wapv5Y9peBvlbsDSgqtJABfK8KQq/jK\nYl3h5elIa1I3uNfczeHOnf1enLOUOlq630yeM/yHizz99G1g+z/guMh5+x/OHraW\niA==\n=+Gxh\n-----END PGP PUBLIC KEY BLOCK-----\n"
	googleSdkGpg = "-----BEGIN PGP PUBLIC KEY BLOCK-----\n\n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\nICgvL2RlcG90L2dvb2dsZTMvcHJvZHVjdGlvbi9ib3JnL2Nsb3VkLXJhcHR1cmUv\na2V5cy9jbG91ZC1yYXB0dXJlLXB1YmtleXMvY2xvdWQtcmFwdHVyZS1zaWduaW5n\nLWtleS0yMDIwLTEyLTAzLTE2XzA4XzA1LnB1YikgPGdsaW51eC10ZWFtQGdvb2ds\nZS5jb20+iQEoBBMBCAAcBQJfyX5eCRCLV8XCg29L6wIbAwUJA8MKsAIZAQAAQXoI\nAJp6SujppQkl3eZW2u0s457BXx63WN6nl7Cak6t7D+lqNFUF4htXVrBWfT/wRx6c\nV6OY/pYVx0Jai4NUhL9CZEVBX03phc3w+LBqWtFXdtMoUEeTBrVwQfgBp8IajRwO\n8lvw3cmcz9GZvUpIYqvtlemaQu1hWqvoeDOX6yeuUfRfCyXmvdLK+oGC9FhNApmk\nqqYER3W52J5WgtIX3zyQp40OMVr5DVKNk8Zb5e6SATyQFRtIPHxjyroKzOHdVVj7\nklR6qKZGAnYANjb/eUS0Bk62bBwGWFjRw9ZWqBrXgTuRMZJYoLpb3o1/L0rsvhim\nnydMwSu6zXzVqrDA18yqeme5AQ0EX8l+XgEIAMNKH4p0F42MhyVVWdEq3RBPnn3f\nIXhZVm39OA25oRnm8qSeihVUir5wkh9j/eVSqrxN7h0SOh20XIp7le56CjqqPOzx\nNgV0IAdeldGhSiVPXib6qQnyuFKuk3RPEiNYlsl2bc2WuBNqblEonDdKHKk1THKi\nWcd6VfGrIoMKTEFIOZHrFx1+3fB1CXkiQgDDF7qfLCZJS2md0YtEl0mtomJpYXYg\nGMU2De9MRiPh6h6Fl3heu8TItWsxspnmbp/WjpBRG7NKtQf0lP+K9moKaE+miwKV\nddfRtuL0rRn/65XKTq/vqazc7nSsirQvRh6Ezwb76qYY7jbX7SoIuR0Pn+MAEQEA\nAYkBHwQYAQgAEwUCX8l+XgkQi1fFwoNvS+sCGwwAABsaCADDuWFJYSVmkZXVIjZv\nYrbo8H3d9UjW1BhXftYgSUImxqMX79kN8knZ3qdy5BKV4flZFxxziFl8rDKwD96m\nqsKfBzvIqxOhnizDzGwkXVS1DkQ22R0CtWIEkdLhjZVnkMRSavkSCEa1kWH1rpou\nh4xUz+etkIuy0NWg8oSkuiQe7+B5sLB9P+f1TV2iUUii+Dp4oAKXnE7H6vNwBWnt\nHg/nbuN9+53Pd0ox0uE+uy0bTDlhvJ88/BYw4+n6RQXjl5h1HuQoqTT1Wsk2QQYn\nQpBDqKapHVOfSTK7hy+R+qcJQXSuufsJ0GotIK/VBGGCvpMKFB/vKdaw/PmAZ1WO\nHe0jmQENBGCRt7MBCADkYJHHQQoL6tKrW/LbmfR9ljz7ib2aWno4JO3VKQvLwjyU\nMPpq/SXXMOnx8jXwgWizpPxQYDRJ0SQXS9ULJ1hXRL/OgMnZAYvYDeV2jBnKsAIE\ndiG/e1qm8P4W9qpWJc+hNq7FOT13RzGWRx57SdLWSXo0KeY38r9lvjjOmT/cuOcm\njwlDT9XYf/RSO+yJ/AsyMdAr+ZbDeQUd9HYJiPdI04lGaGM02MjDMnx+monc+y54\nt+Z+ry1WtQdzoQt9dHlIPlV1tR+xV5DHHsejCZxu9TWzzSlL5wfBBeEz7R/OIziv\nGJpWQdJzd+2QDXSRg9q2XYWP5ZVtSgjVVJjNlb6ZABEBAAG0VEFydGlmYWN0IFJl\nZ2lzdHJ5IFJlcG9zaXRvcnkgU2lnbmVyIDxhcnRpZmFjdC1yZWdpc3RyeS1yZXBv\nc2l0b3J5LXNpZ25lckBnb29nbGUuY29tPokBTgQTAQoAOBYhBDW6oLM+nrOW9Zyo\nOMC6XObcYxWjBQJgkbezAhsDBQsJCAcCBhUKCQgLAgQWAgMBAh4BAheAAAoJEMC6\nXObcYxWj+igIAMFh6DrAYMeq9sbZ1ZG6oAMrinUheGQbEqe76nIDQNsZnhDwZ2wW\nqgVC7DgOMqlhQmOmzm7M6Nzmq2dvPwq3xC2OeI9fQyzjT72deBTzLP7PJok9PJFO\nMdLfILSsUnmMsheQt4DUO0jYAX2KUuWOIXXJaZ319QyoRNBPYa5qz7qXS7wHLOY8\n9IDqfHt6Aud8ER5zhyOyhytcYMeaGC1g1IKWmgewnhEq02FantMJGlmmFi2eA0EP\nD02GC3742QGqRxLwjWsm5/TpyuU24EYKRGCRm7QdVIo3ugFSetKrn0byOxWGBvtu\n4fH8XWvZkRT+u+yzH1s5yFYBqc2JTrrJvRU=\n=0lhn\n-----END PGP PUBLIC KEY BLOCK-----\n-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: GnuPG v1\n\nmQENBFWKtqgBCADmKQWYQF9YoPxLEQZ5XA6DFVg9ZHG4HIuehsSJETMPQ+W9K5c5\nUs5assCZBjG/k5i62SmWb09eHtWsbbEgexURBWJ7IxA8kM3kpTo7bx+LqySDsSC3\n/8JRkiyibVV0dDNv/EzRQsGDxmk5Xl8SbQJ/C2ECSUT2ok225f079m2VJsUGHG+5\nRpyHHgoMaRNedYP8ksYBPSD6sA3Xqpsh/0cF4sm8QtmsxkBmCCIjBa0B0LybDtdX\nXIq5kPJsIrC2zvERIPm1ez/9FyGmZKEFnBGeFC45z5U//pHdB1z03dYKGrKdDpID\n17kNbC5wl24k/IeYyTY9IutMXvuNbVSXaVtRABEBAAG0Okdvb2dsZSBDbG91ZCBQ\nYWNrYWdlcyBSUE0gU2lnbmluZyBLZXkgPGdjLXRlYW1AZ29vZ2xlLmNvbT6JATgE\nEwECACIFAlWKtqgCGy8GCwkIBwMCBhUIAgkKCwQWAgMBAh4BAheAAAoJEPCcOUw+\nG6jV+QwH/0wRH+XovIwLGfkg6kYLEvNPvOIYNQWnrT6zZ+XcV47WkJ+i5SR+QpUI\nudMSWVf4nkv+XVHruxydafRIeocaXY0E8EuIHGBSB2KR3HxG6JbgUiWlCVRNt4Qd\n6udC6Ep7maKEIpO40M8UHRuKrp4iLGIhPm3ELGO6uc8rks8qOBMH4ozU+3PB9a0b\nGnPBEsZdOBI1phyftLyyuEvG8PeUYD+uzSx8jp9xbMg66gQRMP9XGzcCkD+b8w1o\n7v3J3juKKpgvx5Lqwvwv2ywqn/Wr5d5OBCHEw8KtU/tfxycz/oo6XUIshgEbS/+P\n6yKDuYhRp6qxrYXjmAszIT25cftb4d4=\n=/PbX\n-----END PGP PUBLIC KEY BLOCK-----"
)

func TestDistributionFile_Architecture(t *testing.T) {
	adr, err := LoadDistroRegistry("../../distributions")
	require.NoError(t, err)
	d, err := adr.Available(false).Get("centos-8")
	require.NoError(t, err)

	arch, err := d.Architecture("x86_64")
	require.NoError(t, err)

	// don't test packages, they are huge
	arch.Packages = nil

	require.Equal(t, &Architecture{
		ImageTypes: []string{"aws", "gcp", "azure", "ami", "vhd", "guest-image", "image-installer", "oci", "vsphere", "vsphere-ova", "wsl"},
		Repositories: []Repository{
			{
				Id:       "baseos",
				Baseurl:  common.ToPtr("http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/"),
				Rhsm:     false,
				CheckGpg: common.ToPtr(true),
				GpgKey:   common.ToPtr(centosGpg),
			},
			{
				Id:       "appstream",
				Baseurl:  common.ToPtr("http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/"),
				Rhsm:     false,
				CheckGpg: common.ToPtr(true),
				GpgKey:   common.ToPtr(centosGpg),
			},
			{
				Id:       "extras",
				Baseurl:  common.ToPtr("http://mirror.centos.org/centos/8-stream/extras/x86_64/os/"),
				Rhsm:     false,
				CheckGpg: common.ToPtr(true),
				GpgKey:   common.ToPtr(centosGpg),
			},
			{
				Id:            "google-compute-engine",
				Baseurl:       common.ToPtr("https://packages.cloud.google.com/yum/repos/google-compute-engine-el8-x86_64-stable"),
				Rhsm:          false,
				CheckGpg:      common.ToPtr(true),
				GpgKey:        common.ToPtr(googleSdkGpg),
				ImageTypeTags: []string{"gcp"},
			},
			{
				Id:            "google-cloud-sdk",
				Baseurl:       common.ToPtr("https://packages.cloud.google.com/yum/repos/cloud-sdk-el8-x86_64"),
				Rhsm:          false,
				CheckGpg:      common.ToPtr(true),
				GpgKey:        common.ToPtr(googleSdkGpg),
				ImageTypeTags: []string{"gcp"},
			},
		},
	}, arch,
	)

	arch, err = d.Architecture("unsupported")
	require.Nil(t, arch)
	require.Error(t, err, "Architecture not supported")
}

func TestArchitecture_FindPackages(t *testing.T) {
	adr, err := LoadDistroRegistry("../../distributions")
	require.NoError(t, err)
	d, err := adr.Available(false).Get("centos-8")
	require.NoError(t, err)

	arch, err := d.Architecture("x86_64")
	require.NoError(t, err)

	pkgs := arch.FindPackages("vim")
	require.ElementsMatch(t, []Package{
		{
			Name:    "vim-minimal",
			Summary: "A minimal version of the VIM editor",
		},
		{
			Name:    "vim-common",
			Summary: "The common files needed by any version of the VIM editor",
		},
		{
			Name:    "vim-enhanced",
			Summary: "A version of the VIM editor which includes recent enhancements",
		},
		{
			Name:    "vim-X11",
			Summary: "The VIM version of the vi editor for the X Window System - GVim",
		},
		{
			Name:    "vim-filesystem",
			Summary: "VIM filesystem layout",
		},
	}, pkgs)

	d, err = adr.Available(true).Get("rhel-84")
	require.NoError(t, err)

	arch, err = d.Architecture("x86_64")
	require.NoError(t, err)

	pkgs = arch.FindPackages("vim")
	require.ElementsMatch(t, []Package{
		{
			Name:    "vim-minimal",
			Summary: "A minimal version of the VIM editor",
		},
		{
			Name:    "vim-common",
			Summary: "The common files needed by any version of the VIM editor",
		},
		{
			Name:    "vim-enhanced",
			Summary: "A version of the VIM editor which includes recent enhancements",
		},
		{
			Name:    "vim-X11",
			Summary: "The VIM version of the vi editor for the X Window System - GVim",
		},
		{
			Name:    "vim-filesystem",
			Summary: "VIM filesystem layout",
		},
	}, pkgs)

	// load the test distributions and check that a distro with no_package_list == true works
	adr, err = LoadDistroRegistry("testdata/distributions")
	require.NoError(t, err)

	d, err = adr.Available(true).Get("no-packages-distro")
	require.NoError(t, err)

	arch, err = d.Architecture("x86_64")
	require.NoError(t, err)

	pkgs = arch.FindPackages("vim")
	require.Nil(t, pkgs)

}

func TestInvalidDistribution(t *testing.T) {
	_, err := readDistribution("../../distributions", "none")
	require.Error(t, err, DistributionNotFound)
}

func TestDistributionFileIsRestricted(t *testing.T) {
	distsDir := "testdata/distributions"

	t.Run("distro is not restricted, has no restricted_access field", func(t *testing.T) {
		d, err := readDistribution(distsDir, "rhel-90")
		require.NoError(t, err)
		actual := d.IsRestricted()
		expected := false
		require.Equal(t, expected, actual)
	})

	t.Run("distro is not restricted, restricted_access field is false", func(t *testing.T) {
		d, err := readDistribution(distsDir, "centos-9")
		require.NoError(t, err)
		actual := d.IsRestricted()
		expected := false
		require.Equal(t, expected, actual)
	})

	t.Run("distro is restricted, restricted_access field is true", func(t *testing.T) {
		d, err := readDistribution(distsDir, "centos-8")
		require.NoError(t, err)
		actual := d.IsRestricted()
		expected := true
		require.Equal(t, expected, actual)
	})
}

func TestArchitecture_validate(t *testing.T) {
	tests := []struct {
		name string
		arch Architecture
		err  error
	}{
		{
			"good",
			Architecture{
				ImageTypes: nil,
				Repositories: []Repository{
					{Baseurl: common.ToPtr("http://example.com/repo1")},
					{Metalink: common.ToPtr("http://example.com/repo2")},
				},
				Packages: nil,
			},
			nil,
		},
		{
			"multiple-sources",
			Architecture{
				ImageTypes: nil,
				Repositories: []Repository{
					{
						Baseurl:  common.ToPtr("http://example.com/repo1"),
						Metalink: common.ToPtr("http://example.com/repo2"),
					},
				},
				Packages: nil,
			},
			RepoSourceError,
		},
		{
			"no-source",
			Architecture{
				ImageTypes: nil,
				Repositories: []Repository{
					{},
				},
				Packages: nil,
			},
			RepoSourceError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.arch.validate()
			require.Equal(t, tt.err, err)
		})
	}
}
