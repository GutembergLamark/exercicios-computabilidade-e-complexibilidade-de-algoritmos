class AFN:
    def __init__(self):
        self.states = {'q0', 'q1'}
        self.start_state = 'q0'
        self.accept_states = {'q1'}
    
    def is_accepted(self, string):
        current_states = {self.start_state}  
        
        for char in string:
            next_states = set()
            for state in current_states:
                if state == 'q0':
                    if char == '0':
                        next_states.add('q1')  
                    else:
                        next_states.add('q0')  
                elif state == 'q1':
                    next_states.add('q1')  
            current_states = next_states
        
        return any(state in self.accept_states for state in current_states)

if __name__ == '__main__':
    afn = AFN()
    print(afn.is_accepted("111"))
    print(afn.is_accepted("110"))
    print(afn.is_accepted("001"))
    print(afn.is_accepted("1"))  
    print(afn.is_accepted("0"))  