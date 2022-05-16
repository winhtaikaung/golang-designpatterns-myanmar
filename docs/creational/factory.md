# Factory Pattern

 Factory Pattern ကတော့ Struct တွေ Create လုပ်တဲ့ pattern တွေထဲက တခုအပါအဝင်ဖြစ်ပါတယ်။ Factory လို့နာမည်ပေးရတာကတော့ သူက စက်ရုံတခုလိုပဲ ကျနော်တို့ လိုအပ်တဲ့ ဟာကို သူ့ဟာသူ ဆောက်ပေးသွားပါတယ်။

# Code Structure

### **iGun.go**

```go
package factory

type iGun interface {
	setName(name string)
	setPower(power int)
	GetName() string
	GetPower() int
}

```
---
### **gun.go**

```go
package factory

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) GetName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) GetPower() int {
	return g.power
}


```
---
### **ak47.go**

```go
package factory

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}


```
---
### **hmk.go**

```go
package factory

type hmk struct {
	gun
}

func newHmk() iGun {
	return &hmk{
		gun: gun{
			name:  "HMK gun",
			power: 1,
		},
	}
}

```
---
### **factory.go**

```go
package factory

import "fmt"

func GetGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "hmk" {
		return newHmk(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}


```


## ရှင်းလင်းချက်

- ဒီမှာဆိုရင် Entry က `factory.go` ပါ 

- Factory ကို ကိုယ်လိုချင်တာကို တောင်းရပါတယ် မရှိရင်တော့ သူက nil ကိုReturn ပြန်မှာဖြစ်ပြီး ရှိရင် တော့ ရှိတဲ့ struct ကို ပြန်ပေးမှာပါ

- `hmkFactory, _ := factory.GetGun("hmk")`

- ဒီလိုဆိုရင် Factory Struct ရဲ့ GetGun Method က တောင်းလိုက်တဲ့ `hmk` gun ဆိုတဲ့ Struct ကို instantiate လုပ်ပြီး ထုတ်ပေးမှာဖြစ်ပါတယ်

- ဒီနေရာမှာဆိုရင် gun ဆိုတဲ့ struct ကတော့ မည်သည့် Gun အမျိုးအစားမဆို ယူပြီး Composition လုပ်လို့ရတဲ့ Struct တခုပဲဖြစ်ပါတယ်။ ဆိုလိုတာက `gun` မှာကြေငြာထားတဲ့ Method, Variables တွေကို Compose လုပ်တဲ့ Struct က သီးသန့်ကြေငြာစရာမလိုပဲ သုံးစွဲခွင့်ရှိတဲ့ သဘောပါပဲ။

- gun ကို တော့ `ak47` နဲ့ `hmk` ဆိုတဲ့ struct ၂ခုက ယူပြီး compose လုပ်ထားပါတယ်။

- ဒီနေရာမှာ Factory Pattern ကိုသုံးရင်ပိုကောင်းတာက အခြေခံCompose တူတဲ့ struct တွေကို ထုတ်ပေးတဲ့ နေရာမျိုးမှာဆိုသုံးသင့်ပါတယ်

## UseCase
- ဥပမာ အစာချက်စက်ရုံ `FoodFactory` -> `catFood`,`dogFood` struct တွေကို ထုတ်ပေးမယ် 

- မသုံးသင့်တာ ကတော့

- ဥပမာ လက်နက်စက်ရုံ `GunFactory`->`ak47`,`catFood` ထုတ်တာမျိုးက တော့ မှားတဲ့ usage ဖြစ်ပါတယ်


