/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of what makes up a car
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}

// Profile describes basic details of what makes up a car
type Profile struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Company  string `json:"company"`
	Position string `json:"position"`
}

// Implement
type UserProfile struct {
	Data string `json:"data"`
}

type UserLogin struct {
	Password string `json:"data"`
}

// QueryProfileResult structure used for handling result of query
type QueryProfileResult struct {
	Key    string `json:"Key"`
	Record *UserProfile
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Car
}

// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
		Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
	}

	for i, car := range cars {
		carAsBytes, _ := json.Marshal(car)
		err := ctx.GetStub().PutState("CAR"+strconv.Itoa(i), carAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// CreateCar adds a new car to the world state with given details
func (s *SmartContract) CreateCar(ctx contractapi.TransactionContextInterface, carNumber string, make string, model string, colour string, owner string) error {
	car := Car{
		Make:   make,
		Model:  model,
		Colour: colour,
		Owner:  owner,
	}

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

// UpdateProfile adds a new car to the world state with given details
func (s *SmartContract) UpdateProfile(ctx contractapi.TransactionContextInterface, userId string, from string, to string, company string, position string) error {
	profile := Profile{
		From:     from,
		To:       to,
		Company:  company,
		Position: position,
	}

	carAsBytes, _ := json.Marshal(profile)

	return ctx.GetStub().PutState("profile-" + userId, carAsBytes)
}

// QueryProfile returns the car stored in the world state with given id
func (s *SmartContract) QueryProfile(ctx contractapi.TransactionContextInterface, carNumber string) (*Profile, error) {
	carAsBytes, err := ctx.GetStub().GetState("profile-" + carNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if carAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", carNumber)
	}

	car := new(Profile)
	_ = json.Unmarshal(carAsBytes, car)

	return car, nil
}

// UpdateUserProfile adds a new car to the world state with given details
func (s *SmartContract) UpdateUserProfile(ctx contractapi.TransactionContextInterface, userId string, profile string) error {
	userProfile := UserProfile{
		Data: profile,
	}

	carAsBytes, _ := json.Marshal(userProfile)

	return ctx.GetStub().PutState("profile-" + userId, carAsBytes)
}

// QueryUserProfile returns the car stored in the world state with given id
func (s *SmartContract) QueryUserProfile(ctx contractapi.TransactionContextInterface, userId string) (*UserProfile, error) {
	carAsBytes, err := ctx.GetStub().GetState("profile-" + userId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if carAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", userId)
	}

	car := new(UserProfile)
	_ = json.Unmarshal(carAsBytes, car)

	return car, nil
}

// UpdateUserLogin adds a new car to the world state with given details
func (s *SmartContract) UpdateUserLogin(ctx contractapi.TransactionContextInterface, userName string, password string) error {
	userLogin := UserLogin{
		Password: password,
	}

	infoAsBytes, _ := json.Marshal(userLogin)

	return ctx.GetStub().PutState("login-" + userName, infoAsBytes)
}

// IsValidUser adds a new car to the world state with given details
func (s *SmartContract) IsValidUser(ctx contractapi.TransactionContextInterface, userName string, password string) (isValid bool, err error) {
	userAsBytes, err := ctx.GetStub().GetState("login-" + userName)

	if err != nil {
		return false, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if userAsBytes == nil {
		return false, fmt.Errorf("%s does not exist", userName)
	}

	user := new(UserLogin)
	_ = json.Unmarshal(userAsBytes, user)

	if user.Password == password {
		return true, nil
	}

	return false, nil
}

// UpdateOrgLogin adds a new car to the world state with given details
func (s *SmartContract) UpdateOrgLogin(ctx contractapi.TransactionContextInterface, userName string, password string) error {
	userLogin := UserLogin{
		Password: password,
	}

	infoAsBytes, _ := json.Marshal(userLogin)

	return ctx.GetStub().PutState("login-org-" + userName, infoAsBytes)
}

// IsValidOrg adds a new car to the world state with given details
func (s *SmartContract) IsValidOrg(ctx contractapi.TransactionContextInterface, userName string, password string) (isValid bool, err error) {
	userAsBytes, err := ctx.GetStub().GetState("login-org-" + userName)

	if err != nil {
		return false, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if userAsBytes == nil {
		return false, fmt.Errorf("%s does not exist", userName)
	}

	user := new(UserLogin)
	_ = json.Unmarshal(userAsBytes, user)

	if user.Password == password {
		return true, nil
	}

	return false, nil
}

// QueryAllProfiles returns all profile found in world state
func (s *SmartContract) QueryAllProfiles(ctx contractapi.TransactionContextInterface) ([]QueryProfileResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryProfileResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		if strings.Contains(queryResponse.Key, "profile-"){
			profile := new(UserProfile)
			_ = json.Unmarshal(queryResponse.Value, profile)

			queryResult := QueryProfileResult{Key: queryResponse.Key, Record: profile}
			results = append(results, queryResult)
		}
		
	}

	return results, nil
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryCar(ctx contractapi.TransactionContextInterface, carNumber string) (*Car, error) {
	carAsBytes, err := ctx.GetStub().GetState(carNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if carAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", carNumber)
	}

	car := new(Car)
	_ = json.Unmarshal(carAsBytes, car)

	return car, nil
}

// QueryAllCars returns all cars found in world state
func (s *SmartContract) QueryAllCars(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		car := new(Car)
		_ = json.Unmarshal(queryResponse.Value, car)

		queryResult := QueryResult{Key: queryResponse.Key, Record: car}
		results = append(results, queryResult)
	}

	return results, nil
}

// ChangeCarOwner updates the owner field of car with given id in world state
func (s *SmartContract) ChangeCarOwner(ctx contractapi.TransactionContextInterface, carNumber string, newOwner string) error {
	car, err := s.QueryCar(ctx, carNumber)

	if err != nil {
		return err
	}

	car.Owner = newOwner

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
