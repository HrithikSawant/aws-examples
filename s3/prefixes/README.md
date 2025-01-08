
# S3 Bucket and Prefix Demo

This guide demonstrates how to create an Amazon S3 bucket, create folders using prefixes, and upload files into those folders using shell scripts.
---

### Setup

0. **Make scripts executable**:
Navigate to the `bin` directory and make the scripts executable.
```bash
cd aws-examples/s3/prefixes/bin
chmod u+x *
 ```

## Steps to Demonstrate Prefix Usage in S3:
1. **Create a Random S3 Bucket**
Use the script to create an S3 bucket with a random name. For example:

```bash
./bin/create_bucket demo-bucket-xyz123
```
create an empty S3 bucket

2. **Create a Folder (Prefix) in the Bucket**
Use the script to create a folder (prefix) inside your S3 bucket:

```bash
./bin/create_folder demo-bucket-xyz123 hello/
```
create a folder (prefix)

3. **Upload a File to the Folder (Prefix)**
To upload a file to the hello/ folder, use the script as follows:

```bash
./bin/upload_file demo-bucket-xyz123 hello/ file.txt
```
This command uploads the file.txt file to the hello/ folder inside the bucket.


4. **List the Contents of the Folder**

To verify that the file was uploaded successfully, list the contents of the `hello/` folder:

```sh
aws s3 ls s3://demo-bucket-xyz123/hello/
```

You should see the `file.txt` file listed in the `hello/` folder.

---

### Conclusion

In this tutorial, we demonstrated how to:
1. Create a new S3 bucket.
2. Create folders using prefixes.
3. Upload files to those folders.

You can repeat these steps to organize your files in different "folders" using various prefixes within your S3 bucket.

---

**Note**: Remember that in S3, folders are just a naming convention using prefixes. The actual structure is **flat**, and folders are just part of the key name.
