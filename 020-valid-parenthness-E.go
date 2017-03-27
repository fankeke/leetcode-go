package main

func isValid(s string) bool {
    var temp string
    k:=0
    for _,v:=range(s){
        if v=='(' ||v=='[' || v=='{' {
            temp+=string(v)
            k++
        }else{
            k--
            if k<0 {
                return false
            }
            if v==')'&&temp[k]!='(' {
                return false
            }
            if v==']'&&temp[k]!='['{
                return false
            }
            if v=='}'&&temp[k]!='{' {
                return false
            }
            temp=temp[:len(temp)-1]
        }
    }
    
    return k==0 
    
}


func main(){
	isValid("[()](){}")
}
