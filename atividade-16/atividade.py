
class AFD:
    def __init__(self):
      self.state = 0

    def is_accepted(self, string):
      state = 0  
      for c in string:
          if state == 0:
              if c == '1':
                state = 0
              elif c == '0':
                state = 1
          elif state == 1:
            if c == '1':
              state = 0
            elif c == '0':
              state = 2
          elif state == 2:
              if c == '0':
                state = 2
              elif c == '1':
                pass
      return state == 2  

if __name__ == '__main__':
    afd = AFD()
    print(afd.is_accepted("100011"))   
    print(afd.is_accepted("10101"))   
    print(afd.is_accepted("000"))  
    print(afd.is_accepted("000110"))  
    print(afd.is_accepted("1"))    
    print(afd.is_accepted("100"))    