package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Aveshek Singha.")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Aveshek")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Plz give me GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func createProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Plz give me POST request", 400)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)

	w.WriteHeader(201)

	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
}

func main() {
	mux := http.NewServeMux() //router

	mux.HandleFunc("/hello", helloHandler) // route

	mux.HandleFunc("/about", aboutHandler) // route

	mux.HandleFunc("/products", getProducts)

	mux.HandleFunc("/create-products", createProduct)

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", mux) // "Failed to start the server"
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red. I love orange",
		Price:       100,
		ImgUrl:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green. I hate apple",
		Price:       100,
		ImgUrl:      "https://www.harrisfarm.com.au/cdn/shop/products/40715-done.jpg?v=1623908361&width=1920",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is boring. I feel bored eating banana.",
		Price:       5,
		ImgUrl:      "data:image/jpeg;base64,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",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Grape",
		Description: "Orange is red. I love orange",
		Price:       100,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTArLvUM0zUK6b2w23ErXG52NwAfzZlR2WOtw&s",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "Orange is red. I love orange",
		Price:       100,
		ImgUrl:      "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxMTEhMTExIWFRUXFxgVGBUYFxUXFxUVFRUXFhcXFxcYHSggGB0lHRcVITEhJSkrLi4uFx8zODMtNygtLisBCgoKDg0OGxAQGy0lICUtLS8tLS0tLS0tMDAtLS0tLS03LS0vLS03LS0tLSstLTItLS0tLS0tLS0tLS0tLy0tLf/AABEIALcBEwMBIgACEQEDEQH/xAAbAAEAAwEBAQEAAAAAAAAAAAAAAQIDBAYFB//EADoQAAIBAgIHBQcDAwQDAAAAAAABAgMRITEEBRJBUWFxBiKBkaETMlKxwdHwQmLhFJLxM1Nysgcjgv/EABoBAQACAwEAAAAAAAAAAAAAAAACAwEEBQb/xAAyEQACAgECAgcHAwUAAAAAAAAAAQIDEQQhEjEFEyJBUYGhMmFxkbHR8BRC4QYjcsHx/9oADAMBAAIRAxEAPwD9vsRYuACmz1DRcAFLD88y5FgCQAAAAAAAARIPIkiwBW+RLROJWUkldtJcXgjDeN2CzISPhaw7WUKctlbU5ftWC6uVvS58zSe1lWS/9dNRXFvaf2NC7pTTVc5Z+G/8eoPXqPUiUks3brY/P4a5rt96pJ8r29FgazrbW++RzZ/1BBezD5v/AKZSye3/AKmHxx/uiXi08nfxR4ONX88zdTtbKzuVx/qF98PX+DPCe3aFjxtHWc4e7OVuF7ryZ9fRdf7pxvzX1Ru0dOaezaXZ+n55DhZ9sujHR68Jq8Wn9Oq3Gx14yUlmLyiIABIAAAAiRIAIeRW+RawxAI2ECbcwASAAAAAAAAAAAAAAAAAAZ6RXjBbU5KK4v8xPP6+7WQo3hSSqT3v9MerWb5I8tpOsZVrSnJt7uXSxx9b0vXR2a1xP0Qyfa1t20d9mhTwydSafpD7+R82WsatRPbqOV+OXglgcH9Nm08ee8iE7Kzwdjzeq1t2p9qXl3fILJfTNHbxtfeTo81b6G1KqYabS/VHyW80k89lkuHvN5JNPju4nPSquODw3HTo+iVZWag0v3Wj88zqq6klJK84x5q7f0Loaa17cO354k1XJ7pHJOfdT8zSFVNWurM76OqIrOo30il82xT1JTTbUp9MGvkS/RWY7vmWdTPwPmxrWlay68jaNeztxyfA6NK1JtLCavuurequc9TQKij3o35xxS54YlU9NYllr/ZhwkuaOuhp0qbvHA9Nq3W0aloy7svRvl9jxEJ4OLauuayL0a1rpyV08MUbOi6Qv0r23j3og0mfo4Pj6g1t7VbMvfS/uS3n2D2+nvhfWrIcmVAAFwAAAAAAAAAAAAAAAAAAAAAABxa01nChG8sZfpjvl9lzIWWRri5zeEgbabpkKUHUqSUYrNv5Li+R+Y9ou2NbSJOFJunR4J9+a/c9y/avG5n2i0yelSUqknZe7Be7Hw482fFWhLc2med1XSnXdmG0fqRbOuhZpO7NVRlBXT2o+p8/blHCSwv7x9PRqytxONNNGEa6NpV27HSlt923S2dyuj6tdV7UXsLJyte/Gy3s+5o1GFNWgsd8v1Px+iEdOn2pPC9TbponP4HHoGqp51ZJLcljJ9dy9T61CEYLuq3PN+bOepK5pC5dGyEH/AG4+ff8AnwOhDTRgjoUi8UVpI2ii+PFLmJPBCgI0joikW2SzqSpzOWzIvY6nEynErlCUd0zKlk5dI0aNRW2pQe6UW4+aTszzGs6dSlJQlOd+O3K0lfNep6qZfR6yjKLlFStldJtX3xbyZS7IWSUZvD8fuiFlOVmJ87U2otIt7ZS9nJYwjLuuX/NpYLPNPwPQav1ypS9lWg6Vb4X7s+cJZSR9OlUUkmsmc2stX068Nior701hKL+KL3M9RRo/08P7L+fJ/by88mg2+86weOodoamiVf6fTU3F/wCnpFnszjz5res1zVmb6b2/0OnNRvUmt84QvFY2xu034JmxXfGa32a5p9wUW+R6oGOh6XCrCNSnJShJXUlvX06GxeRAAAAAAAAAABVgE3FyNoRYBYFUcGuNZ+xhdLam/dj4Zvkiu22NcXObwkDi7Vdp6ehxSfeqy9ymv+0uEfnu328BPW8q0nKbvJ57rclyObW0J1KkqtS8pSd3L5JcEuBjo9pJq1+fA8rrtW9V/iuRF5ydVa0lhg938o45ScX3vPcaunKOWK6ZGksVldGgk18DOCiaaxWB3am1S29pu1K+W+T4J7lxf1K6l1c5Sbf+mvn8K+rPQOW5KyWCS3JcC+EFBcUuXcjb02n6x5fIvKolZKyWSXLkQ6hlJ8hE07bZSeTsRikjWCw5HRBHOsNx0U2Sqe5iR00YnTA5YSNIzOhXJI1pJs64WNUYQNYs24s1pIsZzQcx7QjPDCyjGrTOatA7do5qsMzmaqpNbFtcjbVen7F1L3c+jOiv2k0eMNrbvi1spPaus1Z5eJ8hqzPL6+0WXtbQjLvb1azadle73t28jqdGdI2cKql3bfYjdp4yzNGXaTtDWr1HGbUaN9qMUk1hgsbXbt88D62pOwtSdSNTSdmNPN0k25ye5SaSUVxSb8D0up+y9CkoSnCNWqrPbkk9lr4E/ds9+fM+9c71dL9qfM0nZjaJXR6EYRUIRUYxVlFKySW5JGhFxc2CokEXG0ASCu0ACwAABC3kkNAFd4iiX0JAOfTdMhShKc3ZLDm3uS5nitYa0lNylLPFJLctxz9rta+0rWUu5Tuo8HLKUr7+C6cz4lTSZLJ3Xr/J5HpXVy1FnBB9ler8fsZTwa6VXula2HX6HFKCabi7NvwfXgZ1ZXxTxz69SkJt2cVnnw/OZpRWEY5nVDSFinna1uPhvGr9EdWWynaF7yaeCjvx+LcrEaHom1Jxtduyzz68fkel0fR4047Mesnxf2W7+S+mKbb7kXU0uyXuLqUYx2IpKKVksbeO/wATJO+bwXUrUmRE19VfxM7ddaiti/tLMtvKxJTNFy2LUjqi9xojkhMv7QlC3HMw4HbBmkZnPTlc0TN2FngUSiddOobRqHFSZ0Jm3C3KKJxRrGWZKePgYp+ZaEr7h1niVuJoo3ImidvDAqkQnusIwkc1VYc1kV0KnBzpqaT2ZXV90v0vwZ0ziclWFu9k0VUuVV0ZrzLl2ouJ6qwsKFTajGS3pPzVy57dPKyjjtYM2gaAyDNIGhDQBGBBbZABIAAAAAB8XtdrR0NHk4u05vYhycs34JN9bH2j8+/8nVr1NHp/DGU3/wDTSj/1kamut6uiTXPl8weVrVXbc/4ON1d8c96/Mjr0bVs5tKN8ckt/RHptC7EVtm9oRv8AFJ39EzylVM57Vxb+BjB5OjQnKTUsFn18Pqz6Gj081FZtJLNt/Ns9TT7GVrvaqU0uK2n6NL5mi0ejo62ab255Oo9181Bbr8S39Fc97Fwr85IurrcniJxaNoSorGzqP3mso/tj9WUqTL1cyjIXTS7MdkjsU1KEcIzaLRJTROBz5Qzvk2Excg0KbGJh1GVItGPkaQTKsvCQdSTyxk3ps0gzGLLrAmpJLYraNom8ZHPcungWwnjYpkjdEowVS2BpCb5Espvcg0a35kqRntFZP86FnEQwatnLpSuzdMxrLFGJJsnDZnpNXf6UP+KXkdJjocbQiuRse1rWIJe45E/aYABMiAAAAAAAAAAAAGzylXUv9XXlWndQwUOcY4XtzxfiemrRUu5u/V04eP3NEjXuojfhS5L1ByaBqynR9yOOW08ZPxOmrUUVeTSXFny9c67jS7scZ+kev2PMaRrOpUfek35W8kaeo19GkXVxW/guRt06Sdiy9kfW1xrF1HsxbUF4OXXlyPj1GU9sV2zz2o1vWScm92dSulQWEWeVzNsN3KSXM0JWNlyRMnwLwMmyVP8APoYUlxZM42NrBsz9oi91uLeOODGGX2iEyFgSimbyZWxtGRvE5jojuIx54IyRePqapmSZc2YLYpkXuaRZimWuSXiQaNXIhyRmyNoy5McJq5GdLGXVlZTOzUtLaqJ7o4+WXqX6SHWXKJiT4IOR6KEbJLgrFgD2pxQAAAAAAAAACNoKQBIITFwAkYaw0n2dOc/hV/Hd6m+0fL7TY6PO3L5ohY+GDa8CcEnJJng69Zyk23i3nxbJhM53IspYdTw91bc2z0UWsHR7QlSOfaLJmrJbkzbaLoyjldl4kDBb/BGySjTZStfMmo5It4KUlw5mqS6h1LF58RjC2GckKJaKsZxNo/wQWWZBeJCZDkRSwzDNYyLqRgnbPqaRdy9SINGkpl4v/BnFkX5k1Ihg6L/4IvgZOe6/0M51cCUnjcwo5K1Knmz1epdE9nTV/eli+XBHxNQav9pP2kl3Y5fuf2R6o9D0No+CPWyW75Gjrbsvq15gAHdOeAAwARtEdQ5AE3AABVhlrCwBRkv0LWFgCqOfWNPapTX7X6Yo6tkjZMNZWDKeHk/Ka8LSZVM+lr7RPZ1ZLnddHkfNieV1NWJNM7tc8pM0pu5rTRlA2SObOsvUiyJQUSdk13BmcloyyNaiwZjGnbM2jPj5k4csMjL3EPc15cTXZduJnFRW+5Z1PIxLbmYRZZq+/BmlONrLh5GSsWpvmypPGCTWTX7/AIiZb+Hh6FNlErpiZ4nkxgtnbhgTLf6YozLJBTT7jDibSl4q3FWKK6WG/wBES7XWBnUii9vJFIlP0+2BtqzQpVpWyivelw/kau0CVZ2WEf1S+nNnrtF0aNOKjFWS9XxfM7HR3R7txOz2fqauo1KrXDHn9C9GkoxUYqyWCRcA9Qlg5AAAAIkSACJFWSQkuIAswWvyABIAAAAAAAAPPdrtWe0gqkV3o4PnH+PqeE2WmfrbR4ztHqP2bdSC7jzXwv7HM12m4u3HzN7S347L8jz1M2izPYsWizg2QwdJM1RZMzVyWs8TWlHfKJo0vvzEWVRdYFeH3AmWY/OpDY2StrLJIs7l6a4/jEV9izRGVfeOI2T8yG+X59CqwLJElFsjknMnaKSmb6HoFWr7qsvieC/nwLKtNOyWILLMSmorMjndW28+nq3Uk6jUql4w4fql9lz/AMn2NW6kp0rN9+fxPJdFuPpnotH0Qodq7f3HNu1udq/mUo0owioxSSWSRcA7iWNkc8AAAAAAAAAAAAAAAAAAAAAAAAESV8HiiQAeZ1t2aveVL+z7Hmq2iyg7Si0+aP0syr6PCatKKkuaNG/Qws3WxtVaqUdnufmyLJHs9J7OUpe7ePqvU+bW7M1F7sov0OZZ0ZYuW5uR1cGfAsXTO+pqWtH9D8MfkZPV1X/bl/a/saktHYv2lyug+85yFidcdW1P9uXkzop6orP9D8bL5kP0Nsv2+g6+C7zigNg+zS7P1Hm0vG/yPoUNQU17zcvRfc2IdE2y5+pVLWQXI81GHI79G1NVnmtlcXh6ZnqKGjQh7sUui+pqdCnoiuPtvJrT1sn7KPl6HqOlDFrbfPLy+9z6iQB066oVrEFg05TlJ5kwACwiAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAf//Z",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
}
