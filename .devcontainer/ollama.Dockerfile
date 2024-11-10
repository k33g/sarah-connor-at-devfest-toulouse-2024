FROM ollama/ollama:0.3.14
RUN /bin/sh -c "/bin/ollama serve & sleep 1 && ollama pull qwen2.5:0.5b && ollama pull mxbai-embed-large:latest"
ENTRYPOINT ["/bin/ollama"]
EXPOSE 11434
CMD ["serve"]