package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDQ9zoSxIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwPj6vc7E95ExucLuwvUPw+es66WY5K09cvZ95UWhn1ZuejRStdPhdv36dzYPv3u3kUcrrsuCDlxWIiqiQi1RnB6toTIeB24ovn08JXd6fXdr/7d2FhbX7OtcInVR+cps2fPnTv359+09O958QIsjgKdNtyfGjoyH7q7P+BI17zq+vxj1KHg3aortQ4kRQksVw9epN4aJdaZ+fDbSQ9WAUnnuU5Lbu6amX/RPOJUM0914FYvtd+3U1rvV2R25u/dXrZ6a73ch7PtOUtNPv9N/fdCVfvNh6/Ve7Z8Xl95fPbBjZ6KZxmPz5r114uxcspLqdvrdJR7vYy27o3KEP7PY/CxZtYsn53Zn+9173JJ93xuNImN0ULzPNPPZWxPZ8+awi7L1TeJa4X/C6kjHoZP1i0oy/DYNa95U0q7ZmWzgu+XlbHeP94GefZuyBTl/1F0lNfLpvrIieeqNvnsre9n/LnLfP7I+/dr/jqbKwqviVhXa7Bz5/nt3r3hq8WPMljvSnk+g0nm3Vx7r7PzPZ6Xqu9uSVqxkvFb9KzYB1nfzbJmTLfdOPUkR+yLiHnGDPP6gw4nnGuQ0EiKWnCDobGLe2NRA+PFiW4MnC2TFhvN9WrdJ8L06ND799e9V1/7mLxvq+RWjvZfQUsy3d9782zadld/EtcJJ95bnyekXP548j3HpXUV/HV3TOv/ZbB4W6uq+79NyUterie4SPmN+qMbmzbXbTnONfU/t3Vs/uMNm6eeND5yruTYN/2ZU4tYUo7NWDVBYEWWQ+En82M2qad31/ztf7YvJC27es2n5aUxG+bJbu73bzpRybeBI7Om4aVEamLemeotvxZzaa+pLt2653j9zujfIY+6nwYfmmpd+GLK2sTDjpr9IhNm7/FisPP9c+Op5gqLlx09/HV3DubpFv1uN76olyJ4Ln7SCpOS4wI2dSUHE56qZlcb5hnkKa6Zk2SQ1JjPEvqiRybvTVNLj6H9tYrm+3ImXWmbEme9iTPdo9cr4c2ZqKX74Wz4wrVRp09IKJy+cP5I2PH7chuWXg3MLb2bUKS2err2n7NzLUylM0QunUheE71xvfLS+YffnWiovjTj7IsP7Cuywr9qJ+RVPJn3Xnrv0oVvyv5l30vebnnp4f4b14J/Z9x/E3jhq7TDu3fLcl69ZgiWmMQZeMxhUYLajgIOxsyE09Jvdpyavy/fxpbL07n89pN+N4akBUK3Wbebxey7kl21InABm6xz0zP+v14XLr34/6lqgUNLXMc2k1/5e9jLjrz0/F2ssuml7AlB5/BfhzXuWzx63e6cZNKzsVMhqizkynmfi/dOvNk5rTHBbMXVjDM8yaVVHJsNL0nG/LYvNktdcmS9Ye+jzVU7c3L1/XyWXKqp7F27Z8PtmG2/l9V9MVL8k7IkKL55ZYn326QL35kTH8xaceP5mR37rtmL3jwmVl90Qn7fvBW1FefCvPIeH4uzX/+3dP7dtVfs7MrirWf/WG9b+5VJ5vRVGwWXoMfu08Svrbmb9ST2xoMZux78i/2hz8Dw/3+ANztHrrD9fRsWBoZF8gwMsGKIgcEArRhiRxRD4JJn7fudCSDdyGoCvBmZRJgRxRiyyaBiDAa2NYJIvIUawijsToEAAYb/ju0sDFgcxsoGkmdiYGLoZGBgOMoC4gECAAD//7ytS4BkBQAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
