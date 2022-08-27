# go_oop

Kotlin dışındaki programlama dilleri hakkında birşey diyemem ama, Go programlama dilinde nesne yönelimli programlama biraz daha yüzeysel. 
Mesela go programlamada 'kalıtım' kavramı sizi biraz üzebilir. Ya da kotlinde olduğu gibi defaultta bir getter/setter olayı yok. 	:face_with_head_bandage:


### Nesne Kurma
Class yok, bunun yerine c programlamadaki gibi struct kullanılıyor. syntax tam olarak şu şekilde;

``` type ``` *class name* ``` struct ```

örnek olarak; 


```

type Desk struct {
	id     int
	name   string
	active bool
}

```
### Getter/Setter fonksiyonlar
İlk başta da yazılı olduğu üzre getter/setter fonksiyonları biz yazmalıyız.

```

func (this *Desk) GetName() string {
	return this.name
}

func (this *Desk) SetName(name string) {
	this.name = name
}

```
Burada dikkat etmemiz gereken durumlardan ilki fonksiyon nesnenin değil :bangbang: nesne adres göstergecinin eklentisi olmalı.

:warning: Getter/Setter fonksiyonlar paket dışından erişilebilir olmalı. Bu yüzden fonksiyon isimleri *Büyük Harf* ile başlamalı :wink:

### Callback

Geri beslemeler için farklı yollar mevcut.

Bu yöntemlerden birisi ve en pratiği fonksiyonlara parametre olarak callback fonksiyonlar vermek. :nerd_face:

```

func (this *Desk) Open(OnFail func(), OnSuccess func()) {
	if this.active {
		OnFail()
		return
	} else {
		OnSuccess()
		this.active = true
		return
	}
}

```

Ama bu yöntem fonksiyon sayıları arttıkça biraz can sıkıcı olabiliyor. Bu yüzden diğer bir yöntem olarak nesnemize callback değerleri ekleyebiliriz. 
Böylece fonksiyonlar gerektiğinde nesne içerisinden ilgili fonksiyonu çağırabilir :100:

Fonksiyonları nesnelere eklemek birer değişken eklemek gibi. Sadece değişken tipi ```int``` veya ```string``` gibi tipler yerine ```func``` olacak.

örnek -> ``` callback func(bool) ```

:warning: callback değerleri kullanmadan önce mutlaka tanımlanmalı. Fonksiyon değerler primatif tipler gibi default değer almıyor ve bu sebeple program hata veriyor.

##### Fonksiyonları parametre olarak göndermek için iki yöntem var 

1. fonksiyonu dışarıda tanımlayıp parametre olarak fonksiyon ismi göndermek

```

func main() {

	desk := obj.Desk{}
            .
            .
            .
	desk.Open(OnFail, OnSuccess)
}

func OnFail() {
	fmt.Println("open fail")
}

func OnSuccess() {
	fmt.Println("open success")
}

```

2. fonksiyonu doğrudan parametre kısmında *anonim fonksiyon* olarak yazmak
```

	desk.SetCallback(func(flag bool) {
		fmt.Println(flag)
	})
  
```
