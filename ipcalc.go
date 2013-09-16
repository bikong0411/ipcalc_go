package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "math"
)

func main() {
   args := os.Args
   if len(args) != 2 {
      fmt.Printf("Usage:%s ip/mask", args[0])
      os.Exit(1)
   }
   str := args[1]
   arr := strings.Split(str,"/")
   ip, num := arr[0], arr[1]
   number, err := strconv.Atoi(num)
   if err != nil {
      fmt.Println(err)
      os.Exit(2)
   }
   mask := getMask(number)
   aip := strings.Split(ip,".")
   network := getNetwork(aip,mask)
   s_network := strings.Join(itoa(network),".")
   s_mask := strings.Join(itoa(mask),".")
   fmt.Printf(`
     =========output===========
     ip:%s
     mask: %s
     network: %s
   `,ip,s_mask,s_network)
}

func getMask(num int) []int {
    mask := make([]int,4)
    div := num / 8
    mod :=  num % 8
    length := 0
    for i:=1;i<=div;i++ {
        mask[i-1] = 255
        length = length + 1
    }
    if mod != 0 {
       num := 0
       for j:=1;j<=mod;j++ {
          num += int(math.Pow(2,float64(8-j)))
       }
       mask[length] = num
       length = length + 1
    }
    for length <= 3 {
       mask[length] = 0
       length += 1
    }
    return mask
}


func getNetwork(ip []string,netmask []int) [] int {
    network := make([]int,4)
    for k,v := range ip {
        tmp,_ := strconv.Atoi(v)
        network[k] = tmp & netmask[k]
    }
    return network
}

func itoa(arr []int) []string {
    s := make([]string,len(arr))
    for k,v := range arr {
        str := strconv.Itoa(v)
        s[k] = str
    }
    return s
}
