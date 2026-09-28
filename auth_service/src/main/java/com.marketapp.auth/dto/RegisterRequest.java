package com.marketapp.auth.dto;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;

public class RegisterRequest {

    @NotBlank(message = "Email cannot be empty.")
    @Email(message = "Invalid email format")
    private String email;

    @Size(min = 1, message = "First name must be longer than 1 symbol.")
    private String firstName;

    @Size(min = 1, message = "Last name must be longer than 1 symbol.")
    private String lastName;

    @NotBlank(message = "The password cannot be empty.")
    @Size(min = 6, message = "The password must be at least 6 characters long.")
    private String password;

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getFirstName() {
        return firstName;
    }

    public void setFirstName(String firstName) {
        this.firstName = firstName;
    }

    public String getLastName() {
        return lastName;
    }

    public void setLastName(String lastName) {
        this.lastName = lastName;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
}