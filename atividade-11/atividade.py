class AFD:
    def __init__(self):
        self.state = 'q0'  

    def transition(self, char):
        if self.state == 'q0':
            if char == 'a':
                self.state = 'q1'  
            elif char == 'b':
                self.state = 'q0'  
        elif self.state == 'q1':
            if char == 'a':
                self.state = 'q0'  
            elif char == 'b':
                self.state = 'q1'  

    def is_accepted(self, string):
        self.state = 'q0'  
        for char in string:
            self.transition(char)
        return self.state == 'q1'  

if __name__ == '__main__':
    afd = AFD()
    print(afd.is_accepted("a"))       
    print(afd.is_accepted("aa"))      
    print(afd.is_accepted("ab"))     
    print(afd.is_accepted("bba"))     
    print(afd.is_accepted("bb"))      
    print(afd.is_accepted("aaa"))      