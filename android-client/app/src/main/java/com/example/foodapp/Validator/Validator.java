package com.example.foodapp.Validator;

public class Validator {

    public boolean isValidEmail(String email) {
        return email.matches("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$");
    }

    public boolean isValidPassword(String password) {
        return password.matches("^[a-zA-Z0-9!@#$%^&*()_+={}|:<>?,./;'\"\\\\`~]+$");
    }
}
