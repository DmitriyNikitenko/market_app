package com.marketapp.user.dto;

import jakarta.validation.constraints.Size;

public class UpdateProfileRequest {

    @Size(max = 50, message = "The name must not exceed 50 characters")
    private String firstName;

    @Size(max = 50, message = "The surname must not exceed 50 characters")
    private String lastName;

    @Size(max = 20, message = "The phone number must not exceed 20 characters")
    private String phone;

    public UpdateProfileRequest() {
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

    public String getPhone() {
        return phone;
    }

    public void setPhone(String phone) {
        this.phone = phone;
    }
}