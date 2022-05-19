# Singleton Pattern

 Singleton pattern ဆိုတာကတော့ Creational Pattern တွေထဲမှာ အသုံးအများဆုံး ဖြစ်ပါတယ်။ အထူးသဖြင့် Object တခုကို တနေရာထဲမှာပဲ ကြေငြာပြီး Application တခုလုံးမှာခေါ်သုံးချင်တာမျိုးတွေဆို အဓိကသုံးဖြစ်ပါတယ်။ 


# Code Structure

### **singleton.go**

```go
package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func GetSingleTonObject() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creting Single Instance Now")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}

```


## ရှင်းလင်းချက်

- ဒီမှာဆိုရင် `var singleInstance` ဆိုတဲ့ Variable လေးကို Constructor မှာ မထားပဲ အပြင်ထုတ်ထားတာတွေ့မှာပါ ဘာကြောင့်လဲဆိုတော့ static Variable အနေနဲ့ ထားချင်လို့ပါ။

- ပြီးရင်တော့ `GetSingleTonObject()` ဆိုတဲ့ Method မှာ ကတော့ singleInstance ကိုအရင်စစ်ပါတယ် မရှိခဲ့ ရင်တော့ `single` ဆိုတဲ့ struct လေးကို သွားဆောက်ပြီး `singleInstance` ဆိုတဲ့ Variable မှာ သွားပြီးသိမ်းပေးမှာပါ။

- ဒီနေရာမှာ တခုသတိထားစေချင်တာကတော့ `sync.once` ကိုသုံးပြီး `once.Do()` scope မှာ SingleInstance ကို တည်ဆောက်ထားပါတယ်။ အဓိကကတော့ တခါပဲ ဆောက်ချင်တာမို့လို့ပါပဲ။ သူ့အလုပ်လုပ်ပုံက once.Do ကို ထပ်ထပ်ခေါ်ပေမယ့် ပထမတခေါက်ပဲအလုပ်လုပ်ပြီးနောက်အခေါက်တွေမှာ `once`က နောက်တခု မကြေငြာမချင်း အလုပ်မလုပ်ရအောင် လုပ်ထားလို့ပါ။

Ref: https://pkg.go.dev/sync#Once.Do



## UseCase
- ဥပမာအားဖြင့် Database Connection လိုမျိုးတွေ HTTP Client တွေဆိုထပ်ခါထပ်ခါမကြေငြာပဲ သုံးကြတဲ့ နေရာမျိုမှာ Singleton Pattern ကိုသုံးကြပါတယ်။


