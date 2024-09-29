class AFD {
    private String state;
    private boolean rotate;

    public AFD() {
        this.state = "q0";
        this.rotate = false;
    }

    public void transition(char c) {
        switch (state) {
            case "q0":
                if (c == '0') {
                    state = "q1";
                }
                break;
            case "q1":
                if (c == '1') {
                    state = "q2";
                } else {
                    state = "q0";
                }
                break;
            case "q2":
                if (c == '0' && !rotate) {
                    state = "q0";
                    rotate = true;
                } else if (c == '0' && rotate) {
                    state = "q3"; 
                } else if (c == '1') {
                    state = "q0";
                }
                break;
        }
    }

    public boolean isAccepted(String input) {
        state = "q0";
        rotate = false;
        for (char c : input.toCharArray()) {
            transition(c);
        }
        return state.equals("q3");
    }

    
}

public class Main {

   public static void main(String[] args) {
        AFD afd = new AFD();
        System.out.println(afd.isAccepted("010010"));  
        System.out.println(afd.isAccepted("010"));      
        System.out.println(afd.isAccepted("0101010"));  
        System.out.println(afd.isAccepted("101010"));   
        System.out.println(afd.isAccepted("111"));      
        System.out.println(afd.isAccepted("000"));      
    }
}