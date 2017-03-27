//https://leetcode.com/problems/add-binary/?tab=Description
func addBinary(a string, b string) string {
    var (
        s string
        c byte
        )

    idxa,idxb:=len(a)-1,len(b)-1
    for idxa>=0 && idxb>=0 {
        an,bn:=a[idxa],b[idxb]
        cc:=an-'0'+bn-'0'+c
        s,c=string(cc%2+'0')+s,cc/2
        idxa,idxb=idxa-1,idxb-1
    }
    for idxa>=0 {
        cc:=a[idxa]-'0'+c
        s,c=string(cc%2+'0')+s,cc/2
        idxa--
    }
    for idxb>=0 {
        cc:=b[idxb]-'0'+c
        s,c=string(cc%2+'0')+s,cc/2
        idxb--
    }
    if c==1 {
        s="1"+s
    }
    return s
}
