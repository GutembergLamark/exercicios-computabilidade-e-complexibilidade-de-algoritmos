import re

class AFN:
    def is_accepted(self, string):
        padrao = r"110"
        return bool(re.search(padrao, string))

if __name__ == '__main__':
    afn = AFN()
    print(afn.is_accepted("110"))  
    print(afn.is_accepted("101")) 
    print(afn.is_accepted("1110")) 
    print(afn.is_accepted("00110"))
    print(afn.is_accepted("000"))  