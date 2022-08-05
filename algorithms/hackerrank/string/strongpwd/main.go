package main

import "fmt"

func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	var l, u, d, s bool
	l, u, d, u = false, false, false, false
	var count int32

	for _, v := range password {
		if v >= 97 && v <= 122 {
			l = true
		} else if v >= 65 && v <= 90 {
			u = true
		} else if v >= 48 && v <= 57 {
			d = true
		} else if v == 33 || v == 35 || v == 36 || v == 37 ||
			v == 42 || v == 40 || v == 41 || v == 43 ||
			v == 64 || v == 94 || v == 38 || v == 45 {
			s = true
		}
	}
	if s && d && l && u {
		if n >= 6 {
			return 0
		} else {
			return (6 - int32(len(password)))
		}

	}
	if !s {
		count++
	}
	if !l {
		count++
	}
	if !u {
		count++
	}
	if !d {
		count++
	}

	if n+count >= 6 {
		return count
	} else {
		// n + count = 4
		// 6 -4  = 2
		//
		return (6 - (n + count)) + count
	}

}

func main() {
	// fmt.Println(minimumNumber(100, "G0N1+93Gy0C!J4ECIc53+30O9az$K-TgDO^051y2+Qfvt94qI!k)lS(-8g65^A9mt%eRL5WP#f8R)i4O33j#&LNk6H%-pr-@d^*Z"))
	s := "hello my name is Henry";
	temp = s.substring(s.lastIndexOf(" ")+1);
        s = " " +s+ " ";
    for(int i = s.length()-2; i>=0; i--){
        if(s.substring(i, i+1).equals(" ")){
            temp += s.substring(i, s.indexOf(" ", i+1));
        }
    }
    return temp
}
