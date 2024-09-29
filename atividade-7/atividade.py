import re

class AFN:
    def is_accepted(self, string):
      regex = r"^01.*10$"
      return bool(re.match(regex, string))

if __name__ == '__main__':
    afn = AFN()
    print(afn.is_accepted("0110"))   
    print(afn.is_accepted("011010"))  
    print(afn.is_accepted("010"))     
    print(afn.is_accepted("110"))    
    print(afn.is_accepted("0101"))    
    print(afn.is_accepted("0110010")) 
    print(afn.is_accepted("01"))      