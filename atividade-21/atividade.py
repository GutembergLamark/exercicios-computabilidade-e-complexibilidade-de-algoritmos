class AFN:
    def __init__(self):
        self.states = {'q0', 'q1', 'q_reject'}
        self.start_state = 'q0'
        self.accept_states = {'q0'}
    
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
                    if char == 'a':
                        next_states.add('q0')
                    elif char == 'b':
                        next_states.add('q_reject')
                elif state == 'q_reject':
                    next_states.add('q_reject')
            current_states = next_states

        return 'q0' in current_states

if __name__ == '__main__':
    afn = AFN()
    print(afn.is_accepted("a"))       
    print(afn.is_accepted("ab"))      
    print(afn.is_accepted("abb"))     
    print(afn.is_accepted("ababa"))    
    print(afn.is_accepted("bba"))    
    print(afn.is_accepted("ba"))      
    print(afn.is_accepted("b"))       