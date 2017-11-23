# regionCode 
1. read from highRisk.txt.
2. find all the region codes of readed.
3. insert these codes into anther table.
# shared experience
1. slice append 
    
    append is that creates a new array having more memory space 
    and copies the data of original array to the new
    array.**the original slice will not change!!**
```$xslt
func call_s_val()[10][]aa{
	var a [10][]aa
	struct_val(a)
	return a
}

func struct_val(a [10][]aa){
	counter := 1
	for counter < 9{
		var c aa
		c.s = "s"
		c.b = counter
		a[counter] = append(a[counter],c)
		counter++
	}
}
```
2. do not commit when it is not necessary.
```$xslt
func (c config) insertVal(value [10][]BsdHighDangerArea) {
	var wg sync.WaitGroup
	for _, val := range value {
		if len(val) == 0 {
			log.Fatalln("empty insert value")
		}
		wg.Add(1)
		go func(val []BsdHighDangerArea){
			for _, v := range val {
				log.Println("insert table", v)
				dbhelpme.Create(&v)
			//	note:if there is dbhelpme.commit,the insert will fail.
			}
			defer wg.Add(-1)
		}(val)
	}
	wg.Wait()
}
```
3. do not use lock when it is not necessary

Because it will be very slow when program runs.
If using lock,other threads will stop until unlock.

      