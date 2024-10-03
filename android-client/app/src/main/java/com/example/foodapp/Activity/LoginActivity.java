package com.example.foodapp.Activity;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.Toast;

import com.example.foodapp.R;
import com.example.foodapp.Domain.Auth.LoginRequest;
import com.example.foodapp.Domain.Auth.LoginResponse;
import com.example.foodapp.databinding.ActivityLoginBinding;
import com.example.foodapp.Validator.Validator;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class LoginActivity extends BaseActivity {
    ActivityLoginBinding binding;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        binding=ActivityLoginBinding.inflate(getLayoutInflater());
        setContentView(binding.getRoot());

        setVariable();
        getWindow().setStatusBarColor(getResources().getColor(R.color.white));
    }

    private void setVariable() {
        binding.loginButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                String email = binding.emailInp.getText().toString();
                String pass = binding.passInp.getText().toString();
                Validator validator = new Validator();
                if (!validator.isValidEmail(email)) {
                    Toast.makeText(LoginActivity.this, incorrectEmail, Toast.LENGTH_SHORT).show();
                } else if (!validator.isValidPassword(pass)) {
                    Toast.makeText(LoginActivity.this, incorrectPassword, Toast.LENGTH_SHORT).show();
                } else {
                    Login(new LoginRequest(email, pass));
                }
            }
        });

        binding.registerButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                Intent intent = new Intent(LoginActivity.this, RegisterActivity.class);
                startActivity(intent);
            }
        });
    }

    private void Login(LoginRequest req) {
        Call<LoginResponse> call = authApi.Login(req);

        call.enqueue(new Callback<LoginResponse>() {

            @Override
            public void onResponse(Call<LoginResponse> call, Response<LoginResponse> response) {
                if (response.isSuccessful() && response.body() != null) {
                    String token = response.body().getToken();
                    Intent intent = new Intent(LoginActivity.this, DashboardActivity.class);
                    startActivity(intent);
                    finish();
                } else {
                    Toast.makeText(LoginActivity.this, "Неверный адрес электронной почты или пароль", Toast.LENGTH_SHORT).show();
                }
            }

            @Override
            public void onFailure(Call<LoginResponse> call, Throwable t) {
                Toast.makeText(LoginActivity.this, errorConnection, Toast.LENGTH_SHORT).show();
            }
        });
    }
}