package figma

// File contains a Figma file https://www.figma.com/file/:key/:title.
type File struct {
	// A mapping from NodeIDs to component metadata This is to help you
	// determine which components each instance comes from. Currently the only
	// piece of metadata available on components is the name of the component,
	// but more properties will be forthcoming.
	Components struct {
	} `json:"components"`

	// A Node of type DOCUMENT.
	Document      Node `json:"document"`
	SchemaVersion int  `json:"schemaVersion"`
}

// Nodes returns a slice containing all subnodes of a Figma file.
func (f File) Nodes() []Node {
	return flatten(f.Document)
}

func flatten(n Node) []Node {
	nodes := n.Children
	for i := range nodes {
		nodes = append(nodes, flatten(nodes[i])...)
	}
	return nodes
}

// NodeType specifies the kind of node. Different types of nodes have different
// properties.
type NodeType string

const (
	// NodeTypeDocument is the root node
	NodeTypeDocument NodeType = "DOCUMENT"

	// NodeTypeCanvas represents a single page
	NodeTypeCanvas = "CANVAS"

	// NodeTypeFrame is A node of fixed size containing other nodes
	NodeTypeFrame = "FRAME"

	// NodeTypeGroup is a logical grouping of nodes
	NodeTypeGroup = "GROUP"

	// NodeTypeVector is a vector network, consisting of vertices and edges
	NodeTypeVector = "VECTOR"

	// NodeTypeBoolean is a group that has a boolean operation applied to it
	NodeTypeBoolean = "BOOLEAN"

	// NodeTypeStar is a regular star shape
	NodeTypeStar = "STAR"

	// NodeTypeLine is a straight line
	NodeTypeLine = "LINE"

	// NodeTypeEllipse is an ellipse
	NodeTypeEllipse = "ELLIPSE"

	// NodeTypeRegularPolygon is a regular n-sided polygon
	NodeTypeRegularPolygon = "REGULAR_POLYGON"

	// NodeTypeRectangle is a rectangle
	NodeTypeRectangle = "RECTANGLE"

	// NodeTypeText is a text box
	NodeTypeText = "TEXT"

	// NodeTypeSlice is a rectangular region of the canvas that can be exported
	NodeTypeSlice = "SLICE"

	// NodeTypeComponent is a node that can have instances created of it that share the same properties
	NodeTypeComponent = "COMPONENT"

	// NodeTypeInstance is an instance of a component, changes to the component result in the same changes applied to the instance
	NodeTypeInstance = "INSTANCE"
)

// Node contains a group of properties which specifies a leaf in a Figma
// File.
//
// Files in Figma consist of a tree of nodes, each with some number of
// properties. One key property is type, which indicates what kind of node you
// are dealing with. At the root of every file is a DOCUMENT node, which has
// some number of CANVAS nodes as its children. Each canvas node represents a
// Figma Page. This section will explore each node type, what it is, and all the
// properties it has.
type Node struct {
	// A string uniquely identifying this node within the document.
	ID string `json:"id"`

	// The name given to the node by the user in the tool.
	Name string `json:"name"`

	// The type of the node.
	Type NodeType `json:"type"`

	// Whether or not the node is visible on the canvas. (default: true)
	Visible bool `json:"visible"`

	Children        []Node           `json:"children"`
	BackgroundColor Color            `json:"backgroundColor"`
	ExportSettings  []ExportSetting  `json:"exportSettings"`
	BlendMode       BlendMode        `json:"blendMode"`
	PreserveRatio   bool             `json:"preserveRatio"`
	Constraints     LayoutConstraint `json:"constraints"`
	LayoutGrids     []LayoutGrid     `json:"layoutGrids"`

	TransitionNodeID    string `json:"transitionNodeID"`
	Opacity             float64
	AbsoluteBoundingBox Rectangle `json:"absoluteBoundingBox"`
	ClipsContent        bool      `json:"clipsContent,omitempty"`
	Effects             []Effect  `json:"effects"`
	IsMask              bool
	Fills               []Paint     `json:"fills,omitempty"`
	Strokes             []Paint     `json:"strokes,omitempty"`
	StrokeWeight        int         `json:"strokeWeight,omitempty"`
	StrokeAlign         StrokeAlign `json:"strokeAlign,omitempty"`
	Characters          string      `json:"characters,omitempty"`
	Style               TypeStyle   `json:"style,omitempty"`

	// Array with same number of elements as characeters in text box, each
	// element is a reference to the styleOverrideTable defined below and maps
	// to the corresponding character in the characters field. Elements with
	// value 0 have the default type style.
	CharacterStyleOverrides []int `json:"characterStyleOverrides,omitempty"`

	// Map from ID to TypeStyle for looking up style overrides.
	StyleOverrideTable map[int]TypeStyle `json:"styleOverrideTable,omitempty"`

	// ID of component that this instance came from, refers to components table.
	ComponentID string
}
