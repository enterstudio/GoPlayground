# Create a GCP BigQuery Dataset

We are using **Google Cloud Platform's** ***BigQuery** service 
for data storage and retrival.

The service provides **10GB** of free data storage capacity and **1TB** of queries.

This is quite enough for basic data manipulation and analysis workloads 
targetted at Mobile backends and API gateways.

## Objectives

We are going to create a dataset by the name of `my_new_dataset` in the select **GCP** project.

A **dataset** is like a *Database* where multiple tables can be created. 

## Getting Account Credentials

One can't directly access the **BigQuery** outside the application bounds.

This means that we would need special permissions to get access 
to perform *CURD* operations at **BigQuery**.

Typically an application written for **AppEngine** or running 
in **CloudFunctions** can have easy access to **BigQuery**.

But in our case we wish to have access to the same via our Golang program.

For that we need the **Credential file** configuration in the Environment.

Steps to get there:
 1. Go to "API & Service" from **Google Cloud Platform** *Main Menu* in 
 the *Required project* (hamburger menu on top left) 
 2. In the API and Service menu select *Credentials*
 3. Click on *Create Credentials* Button/Dropdown and Select *Service Account Keys*.
 This would take you to the Service Accounts page.
 4. In the *Service account* Dropdown select a previously created credentials 
 or Select *New Service account*
 5. In both cases keep the **JSON** format checked for *Key Type*
 6. If you have to create a *New Service Account* then
    1. Set a understandable *Service Name*
    2. Set the *Role* to **Project > Owner**
    3. Make sure that the *Service Account ID* says `@<PROJECT_ID>.iam.gserviceaccount.com`
    4. Click on *Create* to get the JSON file for Download 
 7. Note that **DO NOT SELECT** `App Engine service account` or `Compute Engine service account` 
 as that would not work. 
 8. After getting the **JSON** from either of the *Credential Accounts* this needs
 to be set into the environment such that our program gets to see it. To do so we 
 to set an *Environment Variable* ***GOOGLE_APPLICATION_CREDENTIALS*** with the path
 to the **JSON** file we got from the previous steps.

 ```shell
 export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/service-accounts-json-file.json"
 ```

 OR:

 ```cmd
 set GOOGLE_APPLICATION_CREDENTIALS=C:\Users\username\Downloads\service-accounts-json-file.json"
 ```
 
 Essentially the *Environment Variable* ***GOOGLE_APPLICATION_CREDENTIALS*** contains the *PATH*
 to the file.

After this configuration we can run our Program.

