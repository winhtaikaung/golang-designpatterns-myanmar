# Builder Pattern

Builder Pattern ကတော့ ရှုပ်ထွေးတဲ့ Objects တွေဆောက်တဲ့ နေရာမှာအဓိကထားသုံးကြတဲ့ Pattern တခုပါပဲ။ ရှုပ်ထွေးတဲ့ Objects ဆိုတာ Object ဆောက်ဖို့ အဆင့် တွေအများကြီးလိုအပ်မယ်,အသွင်တူပေမယ့် လုပ်ဆောင်ပုံခြားနားသွားတဲ့ ရှုပ်ထွေးတဲ့Object တွေ တည်ဆောက်ရာမှာသုံးပါတယ်။

# Code Structure

### **curry.go**

```go
package builder

type curry struct {
	Spices     string
	Meat       string
	Vegetables []string
	Name       string
}

```
---
### **chineseChef.go**

```go
package builder

type chineseChef struct {
	spices     string
	meat       string
	vegetables []string
}

func newChineseChef() *chineseChef {
	return new(chineseChef)
}

func (b *chineseChef) setSpices() {
	b.spices = "Mala Spices"
}

func (b *chineseChef) setMeat() {
	b.meat = "Chicken"
}

func (b *chineseChef) setVegetables() {
	vegetables := []string{"corn", "lili", "bokchoy"}
	b.vegetables = vegetables
}

func (b *chineseChef) getCurry() curry {
	return curry{
		Spices:     b.spices,
		Meat:       b.meat,
		Vegetables: b.vegetables,
		Name:       "mala curry",
	}
}

```
---

### **indianChef.go**

```go
package builder

type indianChef struct {
	spices     string
	meat       string
	vegetables []string
}

func newIndianChef() *indianChef {
	return new(indianChef)
}

func (b *indianChef) setSpices() {
	b.spices = "Tendori Spices"
}

func (b *indianChef) setMeat() {
	b.meat = "Chicken"
}

func (b *indianChef) setVegetables() {
	vegetables := []string{"onion", "garlic", "Tomato"}
	b.vegetables = vegetables
}

func (b *indianChef) getCurry() curry {
	return curry{
		Spices:     b.spices,
		Meat:       b.meat,
		Vegetables: b.vegetables,
		Name:       "Tendori Chicken",
	}
}


```
---

### **chef.go**

```go
package builder

type chef struct {
	builder iChefCooker
}

func NewMainChef(b iChefCooker) *chef {
	mainChef := new(chef)
	mainChef.builder = b
	return mainChef
}

func (d *chef) SetBuilder(b iChefCooker) {
	d.builder = b
}

func (d *chef) MakeCurry() curry {
	d.builder.setMeat()
	d.builder.setSpices()
	d.builder.setVegetables()
	return d.builder.getCurry()
}



```
---

### **builder.go**

```go
package builder

type chef struct {
	builder iChefCooker
}

func NewMainChef(b iChefCooker) *chef {
	mainChef := new(chef)
	mainChef.builder = b
	return mainChef
}

func (d *chef) SetBuilder(b iChefCooker) {
	d.builder = b
}

func (d *chef) MakeCurry() curry {
	d.builder.setMeat()
	d.builder.setSpices()
	d.builder.setVegetables()
	return d.builder.getCurry()
}



```

## ရှင်းလင်းချက်
ကျနော်တို့ Example ဟင်းချက်တာနဲ့ လုပ်ပြထားပါတယ်

လုပ်ရမှာက Chef တွေထဲက တယောက်ကိုခေါ်ပေးပါပြီးရင်သူတို့ ဆီမှာရှိတဲ့ Spice,Vegetables,Meatတွေနဲ့ သူတို့ကျွမ်းကျင်ရာဟင်းတခွက်`curry`ချက်ပေးပါ။


- `builder.go` ကနေစကြည့်ရအောင်

- `GetChefBuilder(str string)` method က ဘာလုပ်ပေးလဲဆိုတော့ ကျနော်တို့ အတွက် လိုအပ်တဲ့ စားဖိုမှုးကို ရွေးထုတ်ပေးပါတယ်
`GetChefBuilder("chinese")` ကိုထည့်ပြီး Run လိုက်တာနဲ့ သူက `chineseChef` struct တခုကို ဆောက်ပေးပါတယ်

- ဒီနေရာမှာ `chineseChef` ဆိုတဲ့ Concrete Builder struct တခုရလာပါပြီ ဒါပေမယ့် ကျနော်တို့ လိုချင်တာက `curry` ဖြစ်ပါတယ် 

- ဟုတ်ပြီဒါဆိုရင်ကျနော်တို့ လွယ်လွယ်လုပ်လို့ရတာ က ရလာတဲ့ `chineseChef` ရဲ့ function တွေအကုန်ခေါ်ပြီး `curry` ကိုတန်းချက်ပေးလို့ရပါတယ်

```go
    chineseChef.setMeat()
    chineseChef.setSpices()
    chineseChef.setVegetables()
    chineseChef.getCurry() //ဒီနေရာမှာဆိုရင်ကျနော်တို့ လိုချင်တဲ့ ဟင်းရပြီ 

```
- ဒီနေရာမှာစဥ်စားလို့ရတာ အကယ်လို့များ နောက်ထပ် အမျိူးအစားမတူတဲ့ Chef တယောက်တိုးလာရင်လဲ ဒီလိုပဲလုပ်ပေးနေရမှာလား? ဟုတ်ပြီဒါဆိုကျနော်တို့ ဒီနေရာမှာ BuilderInterfaceတခု ခံပြီး Chef အမျိူးအစားမတူပေမယ့် လုပ်ဆောင်ပုံတူတာတွေကိုလုပ်ခိုင်းကြည့်ကြမယ်

```go

    <T>.setMeat()
    <T>.setVegetables()
    <T>.setVegetables()
    <T>.getCurry()

```
အဲ့လိုစဥ်းစားမိတဲ့ အတွက် `chef.go` interface ပုံစံလုပ်ဆောင်မယ့် struct ကို ဆောက်လိုက်ပါတယ်
အဲ့ဒီထဲမှာ အပေါ်ကပြောခဲ့သလိုမျိုး ကျနော်တို့ တကယ်ရလာတဲ့ Chef ကိုလက်ခံပြီး chefရဲ့လုပ်ဆောင်နိုင်တဲ့ function များကို တစုတစည်းထည်းစုပြီး လုပ်ဆောင်ကာ `curry` တခုရအောင် `MakeCurry()` function တခု လုပ်လိုက်ပါတယ်။

- နောက်ဆုံးမှာတော့ ဒီလိုမျိုး `curry` ကိုရရှိပါတယ်

```go
    chineseChef := builder.GetChefBuilder("chinese")

	mainChef := builder.NewMainChef(chineseChef)
	curry := mainChef.MakeCurry()
	fmt.Println(curry.Name)

	// Set Builder to indian
	indianChef := builder.GetChefBuilder("indian")
	mainChef.SetBuilder(indianChef)
	curry = mainChef.MakeCurry()
	fmt.Println(curry.Name)
```





