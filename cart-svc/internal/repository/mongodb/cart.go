package mongodb

import (
	"context"
	"errors"
	"github.com/r1nb0/food-app/cart-svc/internal/domain/models"
	"github.com/r1nb0/food-app/cart-svc/internal/repository"
	"github.com/r1nb0/food-app/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type cartRepository struct {
	collection *mongo.Collection
}

func NewCartRepository(collection *mongo.Collection) repository.CartRepository {
	return &cartRepository{
		collection: collection,
	}
}

func (r *cartRepository) GetByID(ctx context.Context, id string) (models.Cart, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Cart{}, err
	}

	var basket models.Cart
	if err := r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&basket); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Cart{}, database.ErrNotFound
		}
		return models.Cart{}, err
	}

	return basket, nil
}

func (r *cartRepository) GetAll(ctx context.Context) ([]models.Cart, error) {
	var carts []models.Cart

	curr, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for curr.Next(ctx) {
		var cart models.Cart
		if err := curr.Decode(&cart); err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}

	return carts, nil
}

func (r *cartRepository) Create(ctx context.Context, cart models.CartCreate) (string, error) {
	if cart.Items == nil {
		cart.Items = make([]models.Item, 0)
	}

	res, err := r.collection.InsertOne(ctx, cart)
	if err != nil {
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID).String()
	return id, nil
}

func (r *cartRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err := r.collection.FindOneAndDelete(ctx, bson.M{"_id": objectID}).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return database.ErrNotFound
		}
		return err
	}

	return nil
}

func (r *cartRepository) AddItem(ctx context.Context, basketID string, item models.Item) error {
	objectBasketID, err := primitive.ObjectIDFromHex(basketID)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateByID(ctx, objectBasketID, bson.D{
		{
			"$push", bson.M{
				"items": item,
			},
		},
	})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return database.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *cartRepository) DeleteItem(ctx context.Context, basketID string, itemID int64) error {
	objectBasketID, err := primitive.ObjectIDFromHex(basketID)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateByID(ctx, objectBasketID, bson.D{
		{
			"$pull", bson.M{
				"items": bson.M{
					"product.id": itemID,
				},
			},
		},
	})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return database.ErrNotFound
		}
		return err
	}

	return nil
}

func (r *cartRepository) UpdateItem(ctx context.Context, basketID string, item models.Item) error {
	panic("implement me")
}
