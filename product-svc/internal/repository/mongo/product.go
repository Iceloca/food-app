package mongo

import (
	"context"
	"errors"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
	"github.com/r1nb0/food-app/product-svc/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//TODO logger

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database, collection string) repository.ProductRepository {
	return &productRepository{
		collection: db.Collection(collection),
	}
}

func (r *productRepository) CreateProduct(ctx context.Context, dto models.CreateProductDTO) (string, error) {
	res, err := r.collection.InsertOne(ctx, &dto)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", repository.ErrAlreadyExists
		}
		return "", err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, dto models.UpdateProductDTO) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(dto.ID)
	if err != nil {
		return false, err
	}

	update := bson.M{}
	if dto.Price != 0 {
		update["price"] = dto.Price
	}
	if dto.Name != "" {
		update["name"] = dto.Name
	}
	if dto.Description != "" {
		update["description"] = dto.Description
	}
	if dto.ImageURL != "" {
		update["imageurl"] = dto.ImageURL
	}
	if dto.Category.Name != "" {
		update["category.name"] = dto.Category.Name
	}
	if dto.Category.ImageURL != "" {
		update["category.imageurl"] = dto.Category.ImageURL
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": update})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, repository.ErrNotFound
		}
		return false, err
	}

	return true, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, id string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, repository.ErrNotFound
		}
		return false, err
	}

	return true, nil
}

func (r *productRepository) GetProductByID(ctx context.Context, id string) (models.Product, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Product{}, err
	}

	res := r.collection.FindOne(ctx, bson.M{"_id": objectID})
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return models.Product{}, repository.ErrNotFound
		}
		return models.Product{}, res.Err()
	}

	var product models.Product
	if err := res.Decode(&product); err != nil {
		return models.Product{}, err
	}
	product.ID = id

	return product, nil
}

func (r *productRepository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	defer func() {
		if _ = cur.Close(ctx); err != nil {
			if err != nil {
				//TODO logger errors.Join(err, errCur)
			}
			//TODO logger errCur
		}
	}()

	var products []models.Product
	for cur.Next(ctx) {
		var product models.Product
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		if objectID, ok := cur.Current.Lookup("_id").ObjectIDOK(); ok {
			product.ID = objectID.Hex()
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) GetProductsByCategory(ctx context.Context, categoryName string) ([]models.Product, error) {
	cur, err := r.collection.Find(ctx, bson.M{"category.name": categoryName})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	defer func() {
		if _ = cur.Close(ctx); err != nil {
			if err != nil {
				//TODO logger errors.Join(err, errCur)
			}
			//TODO logger errCur
		}
	}()

	var products []models.Product
	for cur.Next(ctx) {
		var product models.Product
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		if objectID, ok := cur.Current.Lookup("_id").ObjectIDOK(); ok {
			product.ID = objectID.Hex()
		}
		products = append(products, product)
	}

	return products, nil
}
