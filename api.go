package figma

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiURL = "https://api.figma.com"
)

// Client allows you to interact with the Figma APIs.
type Client struct {
	client *http.Client
	token  string
}

// New returns a client initialized with the personal access token provided.
func New(token string) *Client {
	c := &http.Client{
		Timeout: 1 * time.Minute,
	}

	return &Client{
		client: c,
		token:  token,
	}
}

// File returns the document refered to by key.
//	key is the file to export from.
//
// The file key can be parsed from any Figma file url:
// https://www.figma.com/file/:key/:title.
func (c *Client) File(key string) (File, error) {
	var res File

	path := fmt.Sprintf("%s/v1/files/%s", apiURL, key)
	if err := get(c.client, c.token, path, &res); err != nil {
		return res, err
	}

	return res, nil
}

// Images returns a map of URLs for rendered images of the nodes provided.
//  key is the file to export images from.
//  format specifies the image output format.
//  scale is the image scaling factor, it must be between 0.1 and 4.0.
//  ids is a list of node IDs to render.
//
// Important: the image map may contain values that are null. This indicates
// that rendering of that specific node has failed. This may be due to the node
// id not existing, or other reasons such has the node having no renderable
// components. It is guaranteed that any node that was requested for rendering
// will be represented in this map whether or not the render succeeded.
func (c *Client) Images(key string, scale float64, i ImageFormat, ids ...string) (Images, error) {
	var res imageResponse
	if scale < 0.1 || scale > 4.0 {
		return nil, errors.New("scale must be between 0.1 and 4.0")
	}

	if len(ids) == 0 {
		return nil, errors.New("must provide at least one node")
	}

	v := url.Values{}
	v.Add("scale", fmt.Sprintf("%.1f", scale))
	v.Add("format", string(i))
	v.Add("ids", strings.Join(ids, ","))

	path := fmt.Sprintf("%s/v1/images/%s?%s", apiURL, key, v.Encode())
	if err := get(c.client, c.token, path, &res); err != nil {
		return nil, err
	}

	return res.Images, nil
}

// Comments returns a list of comments made on a file.
//  key is the file to retrieve comments from.
func (c *Client) Comments(key string) (Comments, error) {
	var res commentResponse

	path := fmt.Sprintf("%s/v1/files/%s/comments", apiURL, key)
	if err := get(c.client, c.token, path, &res); err != nil {
		return nil, err
	}

	return res.Comments, nil
}

// AddComment posts a new comment on the file.
//  key is the file to add the comment to.
// 	message is the text contents of the comment to post.
//	v is the absolute canvas position of where to place the comment.
func (c *Client) AddComment(key, message string, v Vector) (Comment, error) {
	var res Comment

	input := map[string]interface{}{
		"message":     message,
		"client_meta": v,
	}

	path := fmt.Sprintf("%s/v1/files/%s/comments", apiURL, key)
	if err := post(c.client, c.token, path, input, &res); err != nil {
		return res, err
	}

	return res, nil
}

// TeamProjects lists the projects for a specified team. Note that this will
// only return projects visible to the authenticated user or owner of the
// developer token.
//	teamID is the id of the team to list projects from
func (c *Client) TeamProjects(teamID string) ([]TeamProject, error) {
	var res teamProjectsResponse

	path := fmt.Sprintf("%s/v1/teams/%s/projects", apiURL, teamID)
	if err := get(c.client, c.token, path, &res); err != nil {
		return nil, err
	}

	return res.Projects, nil
}

func get(c *http.Client, token, url string, res interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("%s: failed to build request: %s", url, err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Figma-Token", token)

	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("%s: failed to peform request: %s", url, err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return fmt.Errorf("%s: failed to decode request: %s", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: expected status %d got %d", url, http.StatusOK, resp.StatusCode)
	}

	return nil
}

func post(c *http.Client, token, url string, body, res interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("%s: failed to marshal body: %s", url, err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("%s: failed to build request: %s", url, err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Figma-Token", token)

	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("%s: failed to peform request: %s", url, err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return fmt.Errorf("%s: failed to decode request: %s", url, err)
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("%s: expected status %d got %d", url, http.StatusOK, resp.StatusCode)
	}

	return nil
}
