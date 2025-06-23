export const dynamic = "force-dynamic";

import { type NextRequest, NextResponse } from "next/server"
import { promises as fs } from "fs"
import path from "path"

export async function POST(request: NextRequest) {
  try {
    console.log("API called: send-to-seeds");
    const { category, vocabulary } = await request.json()
    console.log("Received category:", category);
    console.log("Received vocabulary:", Array.isArray(vocabulary) ? `Array of length ${vocabulary.length}` : vocabulary);

    if (!category || !vocabulary) {
      console.log("Missing category or vocabulary");
      return NextResponse.json({ error: "Category and vocabulary are required" }, { status: 400 })
    }

    // Define the base path and create the category-specific directory
    const basePath = "/mnt/e/valk/free-genai-bootcamp-2025/Week_1/backend_go/seeds"
    const categoryDir = path.join(basePath, category)
    console.log("Category directory:", categoryDir);

    // Create the directory if it doesn't exist
    await fs.mkdir(categoryDir, { recursive: true })
    console.log("Directory created or already exists");

    // Create the JSON file with vocabulary data
    const fileName = `${category.toLowerCase().replace(/\s+/g, "_")}_vocabulary.json`
    const filePath = path.join(categoryDir, fileName)
    console.log("File path:", filePath);

    // Write the vocabulary data to the file
    await fs.writeFile(filePath, JSON.stringify(vocabulary, null, 2), "utf8")
    console.log("Vocabulary file written");

    // Also create a metadata file with generation info
    const metadataFile = path.join(categoryDir, "metadata.json")
    const metadata = {
      category,
      generated_at: new Date().toISOString(),
      item_count: vocabulary.length,
      file_name: fileName,
    }

    await fs.writeFile(metadataFile, JSON.stringify(metadata, null, 2), "utf8")
    console.log("Metadata file written");

    return NextResponse.json({
      success: true,
      path: categoryDir,
      files_created: [fileName, "metadata.json"],
      item_count: vocabulary.length,
    })
  } catch (error) {
    console.error("Error sending to seeds:", error)

    // Handle specific error types
    if (error instanceof Error) {
      if (error.message.includes("ENOENT")) {
        return NextResponse.json(
          {
            error:
              "Base directory path does not exist. Please ensure the path /mnt/e/valk/free-genai-bootcamp-2025/Week_1/backend_go/seeds exists.",
          },
          { status: 400 },
        )
      }
      if (error.message.includes("EACCES")) {
        return NextResponse.json(
          {
            error: "Permission denied. Please check write permissions for the seeds directory.",
          },
          { status: 403 },
        )
      }
    }

    return NextResponse.json({ error: "Failed to send to seeds directory" }, { status: 500 })
  }
}
