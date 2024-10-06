package main

import "fmt"

// Bunun gibi bir durumda iota kullanmak mantıklı fakat, aşağıdaki senaryoda çok iyi olmuyor.
// eğer bunlardan birini kaldıracak olursak iota güncellemek gerekir.
// orta değerlerden birini kaldıracak olursan bir sonraki değeri onun yerine geçmemesi gerekiyor. bu konuda dikkatli olmakta fayfa var.
// hepsini teker teker 1 << 0, 1 << 1, 1 << 2 gibi yazmak daha sağlıklı olacaktır.

type Setting int

const (
	PowerSave Setting = 1 << iota
	BlueTooth
	Wifi
	GPS
	NFC
)

var Settings = map[Setting]string{
	PowerSave: "power save",
	BlueTooth: "bluetooth",
	Wifi:      "wifi",
	GPS:       "gps",
	NFC:       "nfc",
}

func main() {
	fmt.Printf("powerSave: %b, BlueTooth: %b, Wifi: %b, Gps: %b, Nfc: %b \n", PowerSave, BlueTooth, Wifi, GPS, NFC)

	settings := PowerSave

	fmt.Printf("settings: %d, settings byte: %b\n", settings, settings)

	settings = PowerSave | Wifi | GPS | NFC // TODO: aynı zamanda bitlerin sayısal değerlerinin toplamı bunu veriri.
	settingsInt := PowerSave + Wifi + GPS + NFC

	fmt.Printf("settings: %d, settingsInt: %d, settings byte: %b\n", settings, settingsInt, settings)

	settings = deleteSetting(settings, Wifi)

	fmt.Printf("settings: %d, deleted setting: %b, settings byte: %b\n", settings, Wifi, settings)

	isWifiOpen(settings)

	printSettings(settings)
}

func isWifiOpen(settings Setting) {
	fmt.Println(settings & NFC)
}

func deleteSetting(settings, setting Setting) Setting {
	fmt.Printf("and -> %b, %b, %b \n", settings, setting, settings^setting)
	settings &^= setting

	return settings
}

func printSettings(settings Setting) {
	for setting, value := range Settings {
		if settings&setting != 0 {
			fmt.Println("item : ", value)
		}
	}
}
