class AFD:
    def __init__(self):
        self.state = 'q0'  

    def transition(self, char):
        if self.state == 'q0':
            if char == 'a':
                self.state = 'q1'
            elif char == 'b':
                self.state = 'q2'
        elif self.state == 'q1':
            if char == 'a':
                self.state = 'q2'
            elif char == 'b':
                self.state = 'q0'
        elif self.state == 'q2':
            if char == 'a':
                self.state = 'q0'
            elif char == 'b':
                self.state = 'q1'

    def is_accepted(self, string):
        self.state = 'q0'  
        for char in string:
            self.transition(char)
        return self.state == 'q0'  

if __name__ == '__main__':
    afd = AFD()
    print(afd.is_accepted("aaabbb")) 
    print(afd.is_accepted("aab"))     
    print(afd.is_accepted("aabb"))    
    print(afd.is_accepted("aaa"))    
    print(afd.is_accepted("bbbbbb"))     
    print(afd.is_accepted("ababab")) 