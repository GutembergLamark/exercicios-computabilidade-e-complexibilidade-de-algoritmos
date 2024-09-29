class AFD {
    private String state;

    public AFD() {
        this.state = "q0";
    }

    public void transition(char c) {
        switch (state) {
            case "q0":
                if (c == 'a') {
                    state = "q1";
                }
                break;
            case "q1":
                if (c == 'b') {
                    state = "q2";
                } else if (c == 'a') {
                    state = "q1";
                }
                break;
            case "q2":
                if (c == 'a') {
                    state = "q3";
                }
                break;
            case "q3":
                
                break;
        }
    }

    public boolean isAccepted(String input) {
        state = "q0";  
        for (char c : input.toCharArray()) {
            transition(c);
        }
        return state.equals("q2");
    }
}

public class Main {

    public static void main(String[] args) {
        AFD afd = new AFD();
        System.out.println(afd.isAccepted("ab"));      
        System.out.println(afd.isAccepted("aab"));     
        System.out.println(afd.isAccepted("abab"));    
        System.out.println(afd.isAccepted("baab"));    
        System.out.println(afd.isAccepted("bbaab"));   
        System.out.println(afd.isAccepted("baabab"));  
    }
}