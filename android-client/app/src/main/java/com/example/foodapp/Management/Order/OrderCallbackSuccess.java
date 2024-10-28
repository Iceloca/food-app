package com.example.foodapp.Management.Order;

public interface OrderCallback {
    void onSuccess();
    void onError(String message);
}
