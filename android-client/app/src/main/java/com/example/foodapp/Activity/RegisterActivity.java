package com.example.foodapp.Activity;

import android.os.Bundle;
import android.view.View;
import android.widget.Toast;


import com.example.foodapp.Domain.Auth.RegisterRequest;
import com.example.foodapp.Domain.Auth.RegisterResponse;
import com.example.foodapp.databinding.ActivityRegisterBinding;
import com.example.foodapp.Validator.Validator;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class RegisterActivity extends BaseActivity {

    ActivityRegisterBinding binding;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        binding=ActivityRegisterBinding.inflate(getLayoutInflater());
        setContentView(binding.getRoot());

        setVariable();
    }

    private void setVariable() {
        binding.loginButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                String email = binding.emailInp.getText().toString();
                String pass = binding.passInp.getText().toString();
                String repeatPass = binding.repeatPassInp.getText().toString();
                Validator validator = new Validator();
                if (!validator.isValidEmail(email)) {
                    Toast.makeText(RegisterActivity.this, incorrectEmail, Toast.LENGTH_SHORT).show();
                }
                else if (!validator.isValidPassword(pass)) {
                    Toast.makeText(RegisterActivity.this, incorrectPassword, Toast.LENGTH_SHORT).show();
                }
                else if (pass.equals(repeatPass)) {
                    Register(new RegisterRequest(email, pass));
                } else {
                    Toast.makeText(RegisterActivity.this, "Введённые пароли не совпадают", Toast.LENGTH_SHORT).show();
                }
            }
        });
    }

    private void Register(RegisterRequest req) {
        Call<RegisterResponse> call = authApi.Register(req);

        call.enqueue(new Callback<RegisterResponse>() {
            @Override
            public void onResponse(Call<RegisterResponse> call, Response<RegisterResponse> response) {
                if (response.isSuccessful() && response.body() != null) {
                    Toast.makeText(RegisterActivity.this, "Вы успешно зарегистрированы", Toast.LENGTH_SHORT).show();
                } else {
                    Toast.makeText(RegisterActivity.this, "Аккаунт с введённой электронной почтой уже существует", Toast.LENGTH_SHORT).show();
                }
            }

            @Override
            public void onFailure(Call<RegisterResponse> call, Throwable t) {
                Toast.makeText(RegisterActivity.this, errorConnection, Toast.LENGTH_SHORT).show();
            }
        });
    }
}