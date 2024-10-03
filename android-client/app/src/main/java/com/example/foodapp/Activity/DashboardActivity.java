package com.example.foodapp.Activity;

import android.os.Bundle;
import android.view.View;

import androidx.recyclerview.widget.GridLayoutManager;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.foodapp.Adapter.BestProductAdapter;
import com.example.foodapp.Adapter.CategoryAdapter;
import com.example.foodapp.Domain.Product.CategoryResponse;
import com.example.foodapp.Domain.Product.ProductResponse;
import com.example.foodapp.databinding.ActivityDashboardBinding;

import java.util.ArrayList;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class DashboardActivity extends BaseActivity {
    ActivityDashboardBinding binding;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        binding = ActivityDashboardBinding.inflate(getLayoutInflater());
        setContentView(binding.getRoot());


        GetBestProducts();
        GetCategories();
    }


    private void GetBestProducts() {
        Call<ArrayList<ProductResponse>> call = productApi.GetBestProducts();
        binding.progressBarBestFood.setVisibility(View.VISIBLE);
        call.enqueue(new Callback<ArrayList<ProductResponse>>() {
            @Override
            public void onResponse(Call<ArrayList<ProductResponse>> call, Response<ArrayList<ProductResponse>> response) {
                if (response.isSuccessful() && response.body() != null) {
                    binding.bestProductRecycleView.setLayoutManager(
                            new LinearLayoutManager(DashboardActivity.this,
                                    LinearLayoutManager.HORIZONTAL,
                                    false));
                    RecyclerView.Adapter<BestProductAdapter.viewholder> adapter = new BestProductAdapter(response.body());
                    binding.bestProductRecycleView.setAdapter(adapter);
                    binding.progressBarBestFood.setVisibility(View.GONE);
                }
            }

            @Override
            public void onFailure(Call<ArrayList<ProductResponse>> call, Throwable t) {

            }
        });
    }

    private void GetCategories() {
        Call<ArrayList<CategoryResponse>> call = productApi.GetCategories();
        binding.progressBarCategory.setVisibility(View.VISIBLE);
        call.enqueue(new Callback<ArrayList<CategoryResponse>>() {
            @Override
            public void onResponse(Call<ArrayList<CategoryResponse>> call, Response<ArrayList<CategoryResponse>> response) {
                if (response.isSuccessful() && response.body() != null) {
                    binding.categoryRecycleView.setLayoutManager(new GridLayoutManager(DashboardActivity.this, 4));
                    RecyclerView.Adapter<CategoryAdapter.viewholder> adapter = new CategoryAdapter(response.body());
                    binding.categoryRecycleView.setAdapter(adapter);
                    binding.progressBarCategory.setVisibility(View.GONE);
                }
            }

            @Override
            public void onFailure(Call<ArrayList<CategoryResponse>> call, Throwable t) {

            }
        });
    }
}