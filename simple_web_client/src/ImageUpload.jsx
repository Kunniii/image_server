import React, { useState } from "react";
import axios from "axios";

function ImageUpload() {
  const [file, setFile] = useState(null);
  const [imageUrl, setImageUrl] = useState("");

  const handleFileChange = (event) => {
    setFile(event.target.files[0]);
  };

  const handleUpload = async () => {
    if (!file) return;

    // Convert the file to a base64 string
    const reader = new FileReader();
    reader.onloadend = async () => {
      const base64String = reader.result.replace(
        /^data:image\/\w+;base64,/,
        ""
      );
      const ext = file.name.split(".").pop();

      const payload = {
        base_64: base64String,
        ext: ext,
      };

      try {
        const response = await axios.post(
          "http://localhost:8080/upload",
          payload,
          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        );
        setImageUrl(response.data.url);
      } catch (error) {
        console.error("Error uploading the image:", error);
      }
    };
    reader.readAsDataURL(file);
  };

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h2>Image Upload</h2>
      <input type="file" onChange={handleFileChange} />
      <button onClick={handleUpload}>Upload</button>

      {imageUrl && (
        <div style={{ marginTop: "20px" }}>
          <h3>Uploaded Image:</h3>
          <img
            src={`http://localhost:8080${imageUrl}`}
            alt="Uploaded"
            style={{ maxWidth: "100%" }}
          />
        </div>
      )}
    </div>
  );
}

export default ImageUpload;
