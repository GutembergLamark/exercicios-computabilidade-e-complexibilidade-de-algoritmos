class AFD:
    def __init__(self):
        
        self.state = 'q00'  

    def transition(self, char):
        if self.state == 'q00':
            if char == '0':
                self.state = 'q10'
            elif char == '1':
                self.state = 'q01'
        elif self.state == 'q01':
            if char == '0':
                self.state = 'q11'
            elif char == '1':
                self.state = 'q00'
        elif self.state == 'q10':
            if char == '0':
                self.state = 'q00'
            elif char == '1':
                self.state = 'q11'
        elif self.state == 'q11':
            if char == '0':
                self.state = 'q01'
            elif char == '1':
                self.state = 'q10'

    def is_accepted(self, string):
        self.state = 'q00'  
        for char in string:
            self.transition(char)
        return self.state == 'q11'  

if __name__ == '__main__':
  afd = AFD()
  print(afd.is_accepted("01"))  
  print(afd.is_accepted("1100"))  
  print(afd.is_accepted("10"))  
  print(afd.is_accepted("0000"))  
  print(afd.is_accepted("1111"))  