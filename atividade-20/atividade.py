class AFD:
    def __init__(self):
        self.state = 'q0'
    
    def transition(self, char):
        if self.state == 'q0':
            if char == 'a':
                self.state = 'q1'
        elif self.state == 'q1':
            if char == 'b':
                self.state = 'q2'
            elif char == 'a':
                self.state = 'q1'
        elif self.state == 'q2':
            if char == 'a':
                self.state = 'q3'
        elif self.state == 'q3':
            pass  
    
    def is_accepted(self, string):
        self.state = 'q0'  
        for char in string:
            self.transition(char)
        return self.state == 'q2'

if __name__ == '__main__':
    afd = AFD()
    print(afd.is_accepted("ab"))      
    print(afd.is_accepted("aab"))     
    print(afd.is_accepted("abab"))    
    print(afd.is_accepted("baab"))    
    print(afd.is_accepted("bbaab"))   
    print(afd.is_accepted("baabab"))  
    print(afd.is_accepted("aaaaa"))  
    print(afd.is_accepted("bbbbb"))  