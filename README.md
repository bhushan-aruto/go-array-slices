

## 📘 Go Arrays & Slices: Master Guide with Deep Dive Examples

This repository is your all-in-one resource to **understand**, **visualize**, and **master** Arrays and Slices in Go — with code samples, definitions, and practical insights.

---

## 📚 What’s Inside?

This guide helps you:

* 📌 Understand **what arrays and slices are**, with real code
* 📌 Learn how Go handles **memory under the hood**
* 📌 Avoid common **bugs** with slice memory sharing
* 📌 Apply these concepts to **problem-solving and interviews**

---

## 🧠 Core Definitions & Concepts

### 🔸 What is an Array in Go?

An **Array** in Go is a **fixed-size** collection of elements of the same type.

#### ✅ Key Points:

* Stored **contiguously in memory**
* Cannot change its length after creation
* Assignment creates a **copy** (value type)
* Efficient for small, fixed data blocks

```go
var a [3]int = [3]int{1, 2, 3}
fmt.Println(a[0]) // Accessing element
```

> ℹ️ Use arrays when you know the size of data at compile time and want **strict immutability**.

---

### 🔹 What is a Slice in Go?

A **Slice** is a **dynamic, flexible**, and **reference-based** view over an array.

#### ✅ Key Points:

* Points to an **underlying array**
* Has **length** and **capacity**
* Can grow dynamically using `append()`
* **Shares memory** with original array unless copied

```go
s := []int{1, 2, 3}
s = append(s, 4)
```

> ℹ️ Use slices when you want **flexibility**, dynamic size, and **efficient memory use**.

---

### ⚠️ What is Slice Memory Sharing?

When you create a new slice from an existing one (via slicing), it **points to the same underlying array**.

#### ❗ Modifying one will affect the other:

```go
a := []int{10, 20, 30}
b := a[:2]
b[0] = 99
fmt.Println(a[0]) // Output: 99 → changed!
```

> 🧠 **Why?** Because `a` and `b` share the same array underneath.

---

### ✅ How to Prevent Slice Memory Sharing?

Use the built-in `copy()` function to create a **new independent slice**.

```go
a := []int{10, 20, 30}
b := make([]int, len(a))
copy(b, a)
b[0] = 999
fmt.Println(a[0]) // Output: 10 → original unchanged!
```

> ✔️ Best practice when you want to **preserve data integrity**.

---

## 📁 Folder Structure

Each folder is a complete working example:

| Folder                             | Description                            |
| ---------------------------------- | -------------------------------------- |
| `01_array_basics`                  | Declaring and using arrays             |
| `02_array_copy_behavior`           | Shows how arrays copy by value         |
| `03_slice_basics`                  | Slicing, appending, and basics         |
| `04_slice_memory_sharing`          | Demonstrates shared memory in slices   |
| `05_slice_copy_to_prevent_sharing` | Using `copy()` to break memory sharing |

---

## 🚀 How to Run

Run any example by entering the folder and executing the Go file:

```bash
cd 03_slice_basics
go run main.go
```

---

## 💡 Real-World Use Cases

| Scenario               | Array or Slice?          | Reason                     |
| ---------------------- | ------------------------ | -------------------------- |
| Fixed-size RGB pixels  | Array                    | Size is constant and small |
| Dynamic log collection | Slice                    | Can grow with `append()`   |
| HTTP Request headers   | Slice                    | Flexible key-value mapping |
| 2D matrix operations   | Array or Slice of slices | Depends on mutability      |

---

## 🎯 Interview Questions to Prepare

| Question                                            | Key Concept                                            |
| --------------------------------------------------- | ------------------------------------------------------ |
| What happens when you assign an array to another?   | Value-copy (deep copy)                                 |
| How do slices grow in Go?                           | Via `append()`, doubling capacity if needed            |
| How do you prevent two slices from sharing memory?  | Use `copy()`                                           |
| Is slicing an array a deep or shallow copy?         | Shallow — memory is shared                             |
| What is the difference between `len()` and `cap()`? | `len` is number of elements, `cap` is storage capacity |

---

## 📌 Tips

* ✅ Always use `copy()` when modifying a slice derived from another
* ✅ Use `make([]T, len)` for slice creation with size
* ✅ Never assume slices are deep copies unless explicitly copied
* ✅ Use slices for all dynamic or streaming data structures




---

## 📎 Summary: Array vs Slice

| Feature                  | Array          | Slice                 |
| ------------------------ | -------------- | --------------------- |
| Fixed Size               | ✅              | ❌                     |
| Dynamic Growth           | ❌              | ✅                     |
| Memory Shared on Assign  | ❌ (value copy) | ✅ (reference)         |
| Can Use `append()`       | ❌              | ✅                     |
| Copy Required for Safety | ❌              | ✅ (`copy()` function) |

---

