class AFN:
    def __init__(self):
        self.states = {'q0', 'q1', 'q2'}
        self.start_state = 'q0'
        self.accept_states = {'q1', 'q2'}

    def is_accepted(self, string):
        current_states = {self.start_state}  

        for char in string:
            next_states = set()
            for state in current_states:
                if state == 'q0':  
                    if char == 'a':
                        next_states.add('q0')  
                    elif char == 'b':
                        next_states.add('q1')  
                elif state == 'q1':  
                    if char == 'b':
                        next_states.add('q1')  

            current_states = next_states

        
        return 'q1' in current_states or 'q0' in current_states

if __name__ == '__main__':
    afn = AFN()
    print(afn.is_accepted("aaa"))    
    print(afn.is_accepted("aabbb"))  
    print(afn.is_accepted("aaab"))   
    print(afn.is_accepted("aabb"))   
    print(afn.is_accepted("abba"))   
    print(afn.is_accepted("b"))      
    print(afn.is_accepted("ba"))     